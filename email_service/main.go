package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/go-mail/mail"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	pbP "github.com/zjjt/abjnet/product_service/proto/product"
	pbS "github.com/zjjt/abjnet/souscription_service/proto/souscription"
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
			case "souscription.sendmail":
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
	topics := []string{"user.created", "souscription.sendmail"}
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

func sendEmail(from string, to string, cc string, topic string, msghtml string, byteArr []byte) error {
	//first we unpack to slices the to and cc args
	//TO:=strings.Split(to,",")
	//CC:=strings.Split(to,",")

	TO := []string{"thibaut.zehi@groupensia.com"}
	CC := []string{"thibaut.zehi@groupensia.com"}
	m := mail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", TO...)
	for _, v := range CC {
		m.SetAddressHeader("Cc", v, "")
	}

	m.SetHeader("Subject", topic)
	m.SetBody("text/html", msghtml)
	//if there's a file to join first we unmarshall the byteArr into the appropriate format
	if len(byteArr) > 0 {
		var subs *pbS.Souscription
		json.Unmarshal(byteArr, &subs)
	}
	m.Attach("/home/Alex/lolcat.jpg")

	d := mail.NewDialer("smtp.example.com", 587, "user", "123456")
	d.StartTLSPolicy = mail.MandatoryStartTLS

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	// Send the email body.
	log.Println("sendig email to: ", to)
	return nil
}
