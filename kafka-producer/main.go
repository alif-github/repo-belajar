package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"repo-eksperimen-code/src/kafka-producer/producer"
)

func main() {
	var (
		g        *errgroup.Group
		err      error
		writer   = producer.NewKafkaProducer()
		ctx      = context.Background()
		messages = make(chan kafka.Message, 10)
	)

	g, ctx = errgroup.WithContext(ctx)
	g.Go(func() error {
		return catchConsoleMessage(ctx, messages)
	})

	g.Go(func() error {
		return writer.WriteMessage(ctx, messages)
	})

	err = g.Wait()
	if err != nil {
		log.Fatalln(err)
	}
}

func catchConsoleMessage(ctx context.Context, messages chan kafka.Message) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			fmt.Println("input text:")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			input := scanner.Text()
			messages <- kafka.Message{
				Value: []byte(input),
			}
			err := scanner.Err()
			if err != nil {
				return err
			}
		}
	}
}
