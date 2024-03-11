package emqx

import (
	"log"

	"github.com/amirhnajafiz/emqx/internal/model"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func onMessage(channel chan *model.Message) mqtt.MessageHandler {
	return func(c mqtt.Client, message mqtt.Message) {
		channel <- &model.Message{
			Topic:   message.Topic(),
			Payload: message.Payload(),
		}

		message.Ack()
	}
}

func onConnect(c mqtt.Client) {
	rOpts := c.OptionsReader()

	log.Printf("connected to:\n\t%s\n", rOpts.Servers())
}

func onDisconnect(c mqtt.Client, err error) {
	log.Printf("connection lost:\n\t%v\n", err)
}
