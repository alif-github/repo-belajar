package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

func main() {
	network := "tcp"
	address := "localhost:9092"
	topic := "test"
	conn, _ := kafka.DialLeader(context.Background(), network, address, topic, 0)
	_ = conn.SetReadDeadline(time.Now().Add(3 * time.Second))

	batch := conn.ReadBatch(1e3, 1e9)
	bytes := make([]byte, 1e3)
	for {
		_, err := batch.Read(bytes)
		if err != nil {
			break
		}
		fmt.Println(string(bytes))
	}
}
