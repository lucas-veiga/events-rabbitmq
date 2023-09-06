package main

import "github.com/lucas-veiga/events-rabbitmq/pkg/rabbitmq"

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	err = rabbitmq.Publish(ch, "amq.direct", "Hello World!")
}
