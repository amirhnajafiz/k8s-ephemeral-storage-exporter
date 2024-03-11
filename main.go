package main

import (
	"log"

	"github.com/amirhnajafiz/emqx/internal/config"
	"github.com/amirhnajafiz/emqx/internal/emqx"
	"github.com/amirhnajafiz/emqx/internal/model"
)

func handler(channel chan *model.Message) {
	go func() {
		for {
			msg := <-channel
			log.Printf("topic:%s\n\t%s\n", msg.Topic, string(msg.Payload))
		}
	}()
}

func main() {
	// load configurations
	cfg := config.Load()

	// create a channel
	channel := make(chan *model.Message)

	// connect to emqx cluster
	client, err := emqx.New(cfg.EMQX, channel)
	if err != nil {
		log.Fatal(err)
	}

	// create a new handler
	handler(channel)

	// subscribe on rides
	client.Subscribe("rides")

	// publish events
	client.Publish("new-rider", "rides")
}
