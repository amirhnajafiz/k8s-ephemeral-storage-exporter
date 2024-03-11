package main

import (
	"log"
	"time"

	"github.com/amirhnajafiz/emqx/internal/config"
	"github.com/amirhnajafiz/emqx/internal/emqx"
	"github.com/amirhnajafiz/emqx/internal/model"
)

func defaultHandler(channel chan *model.Message) {
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
	defaultHandler(channel)

	// subscribe on rides
	client.Subscribe(cfg.Topic)

	// publish events in period
	tk := time.NewTicker(time.Duration(cfg.Interval) * time.Millisecond)
	for {
		select {
		case <-tk.C:
			client.Publish(cfg.Message, cfg.Topic)
		}
	}
}
