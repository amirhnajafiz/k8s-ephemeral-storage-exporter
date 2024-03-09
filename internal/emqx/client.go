package emqx

import mqtt "github.com/eclipse/paho.mqtt.golang"

type Client interface {
	Publish(msg, topic string) mqtt.PublishToken
	Subscribe(topic string) mqtt.SubscribeToken
}
