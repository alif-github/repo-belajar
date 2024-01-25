package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		errorS error
		sess   *discordgo.Session
	)

	sess, errorS = discordgo.New("Bot MTE1MjEzNTM2OTY0MzU0MDUxMQ.Gh3jIo.JWXad8szTcWik2xZb5f14KZwMy8eDXQTEO4Rg4")
	if errorS != nil {
		log.Fatal(errorS)
	}

	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Content == "Hello" {
			fmt.Println("Channel -> ", m.ChannelID)
			_, errorS = s.ChannelMessageSend("1158610322614145044", "World! @everyone")
			if errorS != nil {
				log.Fatal(errorS)
			}
		}
	})

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	errorS = sess.Open()
	if errorS != nil {
		log.Fatal(errorS)
	}

	defer func() {
		_ = sess.Close()
	}()

	fmt.Println("Connected To Discord!")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
