package messagebroker

import "github.com/tonnytg/messagebroker/pkg/rabbit"

func MessageBrokerStart() {
	rabbit.InitRabbitMQ()
}
