package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	pbP "github.com/zjjt/abjnet/product_service/proto/product"
	pbU "github.com/zjjt/abjnet/user_service/proto/user"
)

//brokerSuscriberranges over a slice of topics and make the broker suscribe to each
//topic based on its particular details
func brokerSuscriber(topics []string, pubsub broker.Broker) error {
	if len(topics) > 0 {
		for _, v := range topics {
			switch v {
			case "user.created":
				_, err := pubsub.Subscribe(v, func(p broker.Event) error {
					var user *pbU.User
					if err := json.Unmarshal(p.Message().Body, &user); err != nil {
						theerror := fmt.Sprintf("%v --from email_service", err)
						return errors.New(theerror)
					}
					log.Println(user)
					go sendEmail(user.Email, "user.created")
					return nil
				})
				return err
			case "product.deleted":
				_, err := pubsub.Subscribe(v, func(p broker.Event) error {
					var product *pbP.Product
					admin := "thibaut.zehi@groupensia.com"
					if err := json.Unmarshal(p.Message().Body, &product); err != nil {
						theerror := fmt.Sprintf("%v --from email_service", err)
						return errors.New(theerror)
					}
					log.Println(product)
					go sendEmail(admin, "product.deleted")
					return nil
				})
				return err
			}
		}
	} else {
		return fmt.Errorf("cannot suscribe since the topic slice is empty --from email_service")
	}
	return nil

}

func main() {
	//slice of topics to suscribe to
	topics := []string{"user.created", "product.deleted"}
	srv := micro.NewService(micro.Name("abjnet.service.email"))
	srv.Init()
	//get the broker instance
	pubsub := srv.Server().Options().Broker
	if err := pubsub.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := pubsub.Connect(); err != nil {
		log.Fatal(err)
	}
	//Subscribe to messages on the broker

	if err := brokerSuscriber(topics, pubsub); err != nil {
		log.Fatalln(err)
	}
	//run the server
	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}

func sendEmail(email string, topic string) error {
	log.Println("sendig email to: ", email)
	return nil
}
