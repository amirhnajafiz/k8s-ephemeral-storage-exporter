package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/amirhnajafiz/emqx/internal/emqx"
)

type Config struct {
	Topic    string
	Message  string
	Interval int
	EMQX     emqx.Config
}

func Load() Config {
	instance := Config{}

	instance.Topic = os.Getenv("C_TOPIC")
	instance.Message = os.Getenv("C_MESSAGE")
	instance.Interval, _ = strconv.Atoi(os.Getenv("C_INTERVAL"))

	instance.EMQX.Host = os.Getenv("EMQX_HOST")
	instance.EMQX.Protocol = os.Getenv("EMQX_PROTOCOL")
	instance.EMQX.Port, _ = strconv.Atoi(os.Getenv("EMQX_PORT"))
	instance.EMQX.ClientID = os.Getenv("EMQX_CLIENT_ID")
	instance.EMQX.KeepAlive, _ = strconv.Atoi(os.Getenv("EMQX_KEEP_ALIVE"))
	instance.EMQX.PingTimeout, _ = strconv.Atoi(os.Getenv("EMQX_PING_TIMEOUT"))

	fmt.Println(instance)

	return instance
}
