package service

import (
	"asset_management_backend/common/logger"
	"asset_management_backend/infras"
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MessageQueueService struct {
	rabbitMqProvider   *infras.RabbitMQProvider
	LabelTaskQueueName string
}

func (a *MessageQueueService) publish(ctx context.Context, bodyData []byte, queueName string) error {

	err := a.rabbitMqProvider.Channel.PublishWithContext(
		ctx,
		"",        // exchange
		queueName, // routing key (tên hàng đợi)
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         bodyData,
			DeliveryMode: amqp.Persistent,
		},
	)
	if err != nil {
		logger.Error("MessageQueueService", "PublishErr", err)
		return err
	}

	return nil
}

func (a *MessageQueueService) SafePublish(ctx context.Context, bodyData []byte, queueName string) error {

	if err := a.publish(ctx, bodyData, queueName); err != nil {
		logger.Error("MessageQueueService", "PublishErr", err)
		if isChannelClosedError(err) {
			a.Reconnect()
			return a.publish(ctx, bodyData, queueName)
		}
		return err
	}

	return nil
}

func (a *MessageQueueService) Reconnect() error {
	logger.Info("MessageQueueService", "Reconnect")
	return a.rabbitMqProvider.Reconnect()
}

func isChannelClosedError(err error) bool {
    if amqpErr, ok := err.(*amqp.Error); ok {
        return amqpErr.Code == 504 // Code 504 = channel/connection is not open
    }
    return false
}

func NewMessageQueueService() *MessageQueueService {
	rabbitMqProvider := infras.GetRabbitMQProvider()
	return &MessageQueueService{
		rabbitMqProvider:   rabbitMqProvider,
		LabelTaskQueueName: rabbitMqProvider.LabelTaskQueue,
	}
}
