package emqx

import (
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Client interface {
	Publish(msg, topic string) mqtt.PublishToken
	Subscribe(topic string) mqtt.SubscribeToken
}

func New(cfg Config, callback mqtt.MessageHandler) (Client, error) {
	// setting debug and error stdout
	mqtt.DEBUG = log.New(os.Stdout, "emqx", 0)
	mqtt.ERROR = log.New(os.Stdout, "emqx", 0)

	// get mqtt connection options and
	// create a new client
	c := mqtt.NewClient(cfg.OPTs(callback))
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return client{
		cfg:  cfg,
		conn: c,
	}, nil
}

type client struct {
	cfg  Config
	conn mqtt.Client
}

func (c client) Publish(msg, topic string) mqtt.PublishToken {

}

func (c client) Subscribe(topic string) mqtt.SubscribeToken {

}
