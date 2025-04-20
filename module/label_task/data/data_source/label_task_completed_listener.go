package datasource

import (
	"asset_management_backend/common/logger"
	"asset_management_backend/infras"
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/uptrace/bun/driver/pgdriver"
)

type LabelTaskCompletedListener struct {
	dbProvider       *infras.DbProvider
	listener         *pgdriver.Listener
	subscriberChans  map[int]chan *int
	mu               sync.Mutex
}

var (
	globalTaskListener *LabelTaskCompletedListener
	once               sync.Once
)

func GetTaskStream(subscriberID int) (<-chan *int, error) {
	once.Do(func() {
		globalTaskListener = &LabelTaskCompletedListener{
			dbProvider:      infras.GetDbProvider(),
			subscriberChans: make(map[int]chan *int),
		}
	})
	return globalTaskListener.subscribe(subscriberID)
}


func (m *LabelTaskCompletedListener) subscribe(subscriberID int) (<-chan *int, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Trả về channel cũ nếu đã tồn tại
	if ch, ok := m.subscriberChans[subscriberID]; ok {
		return ch, nil
	}

	// Nếu chưa có listener, khởi tạo
	if m.listener == nil {
		listener := pgdriver.NewListener(m.dbProvider.Instance)
		if err := listener.Listen(context.Background(), m.dbProvider.Config.TaskCompletedChannel); err != nil {
			return nil, err
		}
		m.listener = listener
		go m.listenLoop()
	}

	// Tạo channel mới, lưu vào map trước khi unlock
	ch := make(chan *int, 1000)
	m.subscriberChans[subscriberID] = ch
	logger.Info("LabelTaskCompletedListener", "New subscriber registered", "SubscriberID", subscriberID)
	return ch, nil
}

// Unsubscribe removes a subscriber and closes their channel
func Unsubscribe(subscriberID int) {
	if globalTaskListener == nil {
		return
	}
	globalTaskListener.mu.Lock()
	defer globalTaskListener.mu.Unlock()

	if ch, ok := globalTaskListener.subscriberChans[subscriberID]; ok {
		close(ch)
		delete(globalTaskListener.subscriberChans, subscriberID)
		logger.Info("LabelTaskCompletedListener", "Subscriber removed", "SubscriberID", subscriberID)
	}

	if len(globalTaskListener.subscriberChans) == 0 {
		go globalTaskListener.shutdownAfterDelay()
	}
}

// Internal loop that distributes notifications to all subscribers
func (m *LabelTaskCompletedListener) listenLoop() {
	for notification := range m.listener.Channel() {
		var payload struct {
			ID int `json:"id"`
		}
		if err := json.Unmarshal([]byte(notification.Payload), &payload); err != nil {
			logger.Error("LabelTaskCompletedListener", "json.Unmarshal payload err", err)
			continue
		}

		m.mu.Lock()
		for id, ch := range m.subscriberChans {
			select {
			case ch <- &payload.ID:
				logger.Info("LabelTaskCompletedListener", "Notification sent", "SubscriberID", id, "TaskID", payload.ID)
			default:
				logger.Warn("LabelTaskCompletedListener", "Subscriber channel full, skipping", "SubscriberID", id)
			}
		}
		m.mu.Unlock()
	}
	logger.Info("LabelTaskCompletedListener", "listenLoop stopped", "Listener channel closed")
}

// Shutdown after short delay to wait for potential new subscribers
func (m *LabelTaskCompletedListener) shutdownAfterDelay() {
	time.Sleep(5 * time.Second)

	m.mu.Lock()
	defer m.mu.Unlock()

	if len(m.subscriberChans) > 0 {
		return
	}

	if m.listener != nil {
		_ = m.listener.Close()
		m.listener = nil
	}

	for id, ch := range m.subscriberChans {
		close(ch)
		delete(m.subscriberChans, id)
	}

	logger.Info("LabelTaskCompletedListener", "shutdownAfterDelay called", "Listener closed")
}
