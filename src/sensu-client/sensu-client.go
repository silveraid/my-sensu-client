package main

import (
	"log"

	"sensu-common/transport/rabbitmq"
	"sensu"
)

func main() {
	cfg, err := sensu.NewConfigFromFlagSet(sensu.ExtractFlags())

	if err != nil {
		log.Fatal(err.Error())
	}

	t, err := rabbitmq.NewRabbitMQTransport(cfg.RabbitMQURI())

	if err != nil {
		log.Fatal(err.Error())
	}

	client := sensu.NewClient(t, cfg)

	client.Start()
}
