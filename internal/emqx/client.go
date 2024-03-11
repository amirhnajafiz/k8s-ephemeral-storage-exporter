package emqx

import (
	"log"
	"os"

	"github.com/amirhnajafiz/emqx/internal/model"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Client manages the three main methods of EMQX client
// publishing, subscribing, and unsubscribing
type Client interface {
	Publish(msg, topic string) mqtt.Token
	Subscribe(topic string) mqtt.Token
	UnSubscribe(topics ...string) mqtt.Token
}

func New(cfg Config, channel chan *model.Message) (Client, error) {
	// setting debug and error stdout
	mqtt.DEBUG = log.New(os.Stdout, "emqx", 0)
	mqtt.ERROR = log.New(os.Stdout, "emqx", 0)

	// get mqtt connection options and
	opts := cfg.OPTs()
	opts.SetDefaultPublishHandler(onMessage(channel))

	// create a new client
	c := mqtt.NewClient(opts)
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

func (c client) Publish(msg, topic string) mqtt.Token {
	return c.conn.Publish(topic, 0, false, msg)
}

func (c client) Subscribe(topic string) mqtt.Token {
	return c.conn.Subscribe(topic, 0, nil)
}

func (c client) UnSubscribe(topics ...string) mqtt.Token {
	return c.conn.Unsubscribe(topics...)
}
