package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	writer *kafka.Writer
}

// NewKafkaProducer создает новый Kafka продюсер
func NewKafkaProducer(brokers []string, topic string) *KafkaProducer {
	return &KafkaProducer{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(brokers...),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		},
	}
}

// Publish отправляет сообщение в Kafka
func (p *KafkaProducer) Publish(message string) error {
	return p.writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("task_event"),
			Value: []byte(message),
		},
	)
}

// Close закрывает соединение
func (p *KafkaProducer) Close() error {
	return p.writer.Close()
}
