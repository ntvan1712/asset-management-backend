package infras

import (
	"asset_management_backend/app_config"
	"asset_management_backend/common/logger"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

var rabbitMQProvider *RabbitMQProvider
var initRabbitMQOnce sync.Once

type RabbitMQProvider struct {
	Channel                 *amqp.Channel
	LabelTaskQueue          string
	CompletedLabelTaskQueue string
}

func GetRabbitMQProvider() *RabbitMQProvider {
	initRabbitMQOnce.Do(initRabbitMQProvider)
	return rabbitMQProvider
}

func initRabbitMQProvider() {
	rabbitMqConfig := app_config.GetAppConfig().RabbitMqConfig
	rabbitMQProvider = &RabbitMQProvider{
		Channel:                 createRabbitMQChannel(rabbitMqConfig),
		LabelTaskQueue:          rabbitMqConfig.LabelTaskQueue,
		CompletedLabelTaskQueue: rabbitMqConfig.CompletedLabelTaskQueue,
	}
	logger.Info("[RabbitMQInfras] Init RabbitMQProvider")
}

func upsertQueue(queueName string, channel *amqp.Channel) error {
	_, err := channel.QueueDeclare(
		queueName, // tên hàng đợi
		true,      // durable
		false,     // auto-delete
		false,     // exclusive
		false,     // no-wait
		amqp.Table{
			"x-max-priority": 10,
		},
	)
	if err != nil {
		logger.Fatal("[RabbitMQInfras] Failed to Register TaskQueue", err)
	}
	return err
}

func (r *RabbitMQProvider) Reconnect() error {
	rabbitMqConfig := app_config.GetAppConfig().RabbitMqConfig
	connection, err := amqp.DialConfig(rabbitMqConfig.URI, amqp.Config{
		Heartbeat: 0,
	})

	if err != nil {
		logger.Error("[RabbitMQInfras] Failed to load DialConfig RabbitMQ", err)
		return err
	}

	//defer connection.Close()

	channel, err := connection.Channel()
	if err != nil {
		logger.Error("[RabbitMQInfras] Failed to connect to RabbitMQ Channel", err)
		return err
	}
	rabbitMQProvider.Channel = channel
	return nil
}

func createRabbitMQChannel(rabbitMqConfig app_config.RabbitMQConfig) *amqp.Channel {

	connection, err := amqp.DialConfig(rabbitMqConfig.URI, amqp.Config{
		Heartbeat: 0,
	})

	if err != nil {
		logger.Fatal("[RabbitMQInfras] Failed to load DialConfig RabbitMQ", err)
	}

	//defer connection.Close()

	channel, err := connection.Channel()
	if err != nil {
		logger.Fatal("[RabbitMQInfras] Failed to connect to RabbitMQ Channel", err)
	}

	queues := [1]string{
		// rabbitMqConfig.CompletedLabelTaskQueue,
		rabbitMqConfig.LabelTaskQueue,
	}

	for _, queue := range queues {
		err = upsertQueue(queue, channel)
		if err != nil {
			logger.Fatal("[RabbitMQInfras] Failed to Register "+queue, err)
		}
	}

	channel.Qos(1, 0, false)
	if err != nil {
		logger.Fatal("[RabbitMQInfras] Failed to Register CompletedTaskQueue", err)
	}
	return channel
}
