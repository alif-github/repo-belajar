package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"golang.org/x/sync/errgroup"
	"log"
	reader2 "repo-eksperimen-code/kafka/reader"
	writer2 "repo-eksperimen-code/kafka/writer"
)

func main() {
	var (
		g *errgroup.Group
	)

	reader := reader2.NewKafkaReader()
	writer := writer2.NewKafkaWriter()

	ctx := context.Background()
	messages := make(chan kafka.Message, 10)
	messageCommitChan := make(chan kafka.Message, 10)

	g, ctx = errgroup.WithContext(ctx)
	g.Go(func() error {
		return reader.FetchMessage(ctx, messages)
	})

	g.Go(func() error {
		return writer.WriteMessage(ctx, messages, messageCommitChan)
	})

	g.Go(func() error {
		return reader.CommitMessages(ctx, messageCommitChan)
	})

	err := g.Wait()
	if err != nil {
		log.Fatalln(err)
	}
}
