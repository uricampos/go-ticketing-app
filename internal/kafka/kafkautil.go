package kafka

import (
	"os"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)

func brokers() []string {
	env := strings.TrimSpace(os.Getenv("KAFKA_BROKERS"))

	if env == "" {
		return []string{"localhost:9092"}
	}

	parts := strings.Split(env, ",")

	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}

	return parts
}

func NewWriter(topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:         kafka.TCP(brokers()...),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		BatchSize:    50,
		BatchTimeout: 10 * time.Millisecond,
		RequiredAcks: kafka.RequireAll,
		Async:        false,
	}
}

func NewReader(topic, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers(),
		Topic:          topic,
		GroupID:        groupID,
		StartOffset:    kafka.FirstOffset,
		CommitInterval: time.Second,
		MaxBytes:       10e6,
	})
}
