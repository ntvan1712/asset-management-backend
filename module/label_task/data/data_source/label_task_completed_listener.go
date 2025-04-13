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
	dbProvider         *infras.DbProvider
	completedTaskIDsCh chan *int
	subscribers        int
	listener           *pgdriver.Listener
	mu                 sync.Mutex
}

var (
	globalTaskListener *LabelTaskCompletedListener
	once               sync.Once
)

func GetTaskStream() (<-chan *int, error) {
	
	var err error
	once.Do(func() {
		globalTaskListener = &LabelTaskCompletedListener{
			dbProvider: infras.GetDbProvider(),
		}
		
	})
	globalTaskListener.subscribers++
	if globalTaskListener.listener == nil {
		if _, err = globalTaskListener.Subscribe(); err != nil {
			return nil, err
		}
	}

	logger.Info("LabelTaskCompletedListener", "Stream called", "Starting listener", "Subscribers", globalTaskListener.subscribers)
	return globalTaskListener.completedTaskIDsCh, nil
}

func (m *LabelTaskCompletedListener) Subscribe() (<-chan *int, error) {
	logger.Info("LabelTaskCompletedListener", "Subscribe called", "Starting Subscribe")
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.completedTaskIDsCh != nil {
		return m.completedTaskIDsCh, nil
	}

	m.completedTaskIDsCh = make(chan *int, 1000)
	listener := pgdriver.NewListener(m.dbProvider.Instance)
	if err := listener.Listen(context.Background(), m.dbProvider.Config.TaskCompletedChannel); err != nil {
		return nil, err
	}
	m.listener = listener

	go m.listenLoop()

	return m.completedTaskIDsCh, nil
}

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
		ch := m.completedTaskIDsCh
		m.mu.Unlock()

		if ch != nil {
			select {
			case ch <- &payload.ID:
				logger.Info("LabelTaskCompletedListener", "TaskID", payload)
			default:
				logger.Info("LabelTaskCompletedListener", "TaskID from default", payload)
			}
		}
	}
	logger.Info("LabelTaskCompletedListener", "listenLoop stopped, listener channel closed")
}

func Unsubscribe() {
	if (globalTaskListener == nil){
		return
	}
	globalTaskListener.mu.Lock()
	defer globalTaskListener.mu.Unlock()

	globalTaskListener.subscribers--
	logger.Info("LabelTaskCompletedListener", "Subscribers", globalTaskListener.subscribers)
	if globalTaskListener.subscribers <= 0 {
		go globalTaskListener.shutdownAfterDelay()
	}
}

func (m *LabelTaskCompletedListener) shutdownAfterDelay() {
	time.Sleep(5 * time.Second) // chờ xem có subscriber mới không

	m.mu.Lock()
	defer m.mu.Unlock()

	if m.subscribers > 0 {
		return
	}

	if m.listener != nil {
		_ = m.listener.Close()
	}
	if m.completedTaskIDsCh != nil {
		close(m.completedTaskIDsCh)
		m.completedTaskIDsCh = nil
	}
	m.listener = nil
	logger.Info("LabelTaskCompletedListener", "shutdownAfterDelay called", "Close stream")
}
