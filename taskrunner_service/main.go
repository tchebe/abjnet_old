package main

import (
	"fmt"
	"log"
	"os"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/robfig/cron/v3"
)

var topic = []string{"taskrunner.deletesubs", "taskrunner.updatesubs"}

func publishEvent(pubsub broker.Broker, topic string) error {
	//create a broker message
	msg := &broker.Message{}
	//publish the message to the broker
	if err := pubsub.Publish(topic, msg); err != nil {
		theerror := fmt.Sprintf("%v --from taskrunner_service", err)
		log.Printf("[PUB] failed %s\n", theerror)
	}
	return nil
}

func main() {
	service := micro.NewService(micro.Name("abjnet.service.taskrunner"))
	service.Init()
	//get the broker instance
	pubsub := service.Server().Options().Broker
	if err := pubsub.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := pubsub.Connect(); err != nil {
		log.Fatal(err)
	}

	//publishing the event and sending all the subs to the email_service
	job := cron.New()
	job.AddFunc(os.Getenv("MAJSUBAT"), func() {
		log.Println("[PUB]publishing the update subscriptions event ")
		if err := publishEvent(pubsub, topic[1]); err != nil {
			fmt.Println(err)
		}
	})
	job.AddFunc(os.Getenv("DELETESUBSAT"), func() {
		log.Println("[PUB]publishing the delete subscriptions event ")
		if err := publishEvent(pubsub, topic[0]); err != nil {
			fmt.Println(err)
		}
	})
	job.Start()
	if err := service.Run(); err != nil {
		theerror := fmt.Sprintf("%v --from taskrunner_service", err)
		fmt.Println(theerror)
	}
	//to stop the goroutine from exiting
	for {
		select {}
	}
}
