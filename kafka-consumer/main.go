package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"golang.org/x/sync/errgroup"
	"log"
	"repo-eksperimen-code/src/kafka-consumer/consumer"
)

func main() {
	var (
		g                 *errgroup.Group
		err               error
		reader            = consumer.NewKafkaReader()
		ctx               = context.Background()
		messageCommitChan = make(chan kafka.Message, 10)
	)

	g, ctx = errgroup.WithContext(ctx)
	g.Go(func() error {
		return reader.FetchMessage(ctx, messageCommitChan)
	})

	g.Go(func() error {
		return reader.CommitMessage(ctx, messageCommitChan)
	})

	err = g.Wait()
	if err != nil {
		log.Fatalln(err)
	}
}
