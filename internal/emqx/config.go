package emqx

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Config struct {
	Protocol    string `koanf:"protocol"`
	Host        string `koanf:"host"`
	Port        int    `koanf:"port"`
	ClientID    string `koanf:"client_id"`
	KeepAlive   int    `koanf:"keep_alive"`
	PingTimeout int    `koanf:"ping_timeout"`
}

func (c Config) DNS() string {
	return fmt.Sprintf("%s://%s:%d", c.Protocol, c.Host, c.Port)
}

func (c Config) OPTs(h mqtt.MessageHandler) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions().
		AddBroker(c.DNS()).
		SetClientID(c.ClientID)

	opts.SetKeepAlive(time.Duration(c.KeepAlive) * time.Second)
	opts.SetPingTimeout(time.Duration(c.PingTimeout) * time.Second)

	opts.SetDefaultPublishHandler(h)
	opts.SetConnectionLostHandler(onDisconnect)
	opts.SetOnConnectHandler(onConnect)

	return opts
}

func onConnect(c mqtt.Client) {
	rOpts := c.OptionsReader()
	log.Printf("connected to:\n\t%s\n", rOpts.Servers())
}

func onDisconnect(c mqtt.Client, err error) {
	log.Printf("connection lost:\n\t%v\n", err)
}
