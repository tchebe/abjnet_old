package main

import (
	"github.com/micro/go-micro/v2/broker"
)

//event to be sent
const topic = "product.deleted"

type service struct {
	repo         repository
	tokenService Authable
	PubSub       broker.Broker
}
