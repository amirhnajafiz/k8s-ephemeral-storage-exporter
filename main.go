package main

import (
	"log"
	"time"

	"github.com/amirhnajafiz/emqx/internal/config"
	"github.com/amirhnajafiz/emqx/internal/emqx"
	"github.com/amirhnajafiz/emqx/internal/model"
)

// defaultHandler is used to handle input messages from
// subscription channel.
func defaultHandler(channel chan *model.Message) {
	go func() {
		for {
			msg := <-channel
			log.Printf("received message from topic: `%s`\n\t%s\n", msg.Topic, string(msg.Payload))
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
		log.Fatalf("failed to connect to emqx cluster: %v", err)
	}

	// create a new handler
	defaultHandler(channel)

	// subscribe on a topic
	if token := client.Subscribe(cfg.Topic); token.Wait() || token.Error() != nil {
		log.Fatalf("failed to subscribe over `%s`: %v", cfg.Topic, token.Error())
	}

	// publish events in period
	tk := time.NewTicker(time.Duration(cfg.Interval) * time.Millisecond)
	for {
		<-tk.C

		if token := client.Publish(cfg.Message, cfg.Topic); token.Wait() || token.Error() != nil {
			log.Printf("failed to publish over `%s`: %v", cfg.Topic, token.Error())
		}
	}
}
