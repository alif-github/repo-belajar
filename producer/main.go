package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"time"
)

func main() {
	network := "tcp"
	address := "localhost:9092"
	topic := "test"
	conn, _ := kafka.DialLeader(context.Background(), network, address, topic, 0)
	_ = conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, _ = conn.WriteMessages(kafka.Message{Value: []byte("Message from lukman")})
}
