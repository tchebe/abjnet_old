package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/server"
	pb "github.com/zjjt/abjnet/payment_service/proto/payment"
	userproto "github.com/zjjt/abjnet/user_service/proto/user"
)

//AuthWrapper is a higher order function that takes a HandlerFunc
//and returns a function,which takes a context,request and response interface.
//The token is extracted from the context set in our consignement-cli,
//that token is then sent over to the user service to be validated.
//if valid the call is passed along to the handler, if not an error is returned
func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, res interface{}) error {
		// This skips our auth check if DISABLE_AUTH is set to true
		//in order to test the service in isolation
		if os.Getenv("DISABLE_AUTH") == "true" {
			return fn(ctx, req, res)
		}
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request --from payment_service l35")
		}
		//it changed from token to Token
		token := meta["Token"]
		log.Println("Authenticated with token: ", token)
		//Checking if the token is a valid one
		authClient := userproto.NewUserService("abjnet.service.user", client.DefaultClient)
		_, err := authClient.ValidateToken(context.Background(), &userproto.Token{
			Token: token,
		})
		if err != nil {
			theerror := fmt.Sprintf("%v --from payment_service l46", err)
			return errors.New(theerror)
		}
		err = fn(ctx, req, res)
		return err
	}
}

var topic = "taskrunner.deletepayments"

func publishEvent(subs []*pb.Payment, pubsub broker.Broker, topic string) error {
	//when sending an event we have to serialize it to bytes
	//we are sending to our ecosystem the event payment.sendmail with the details
	//with all today's subscription
	body, err := json.Marshal(subs)
	if err != nil {
		theerror := fmt.Sprintf("%v --from payment_service", err)
		return errors.New(theerror)
	}

	//create a broker message
	msg := &broker.Message{
		Header: map[string]string{
			"to":    os.Getenv("TO"),
			"cc":    os.Getenv("CC"),
			"objet": fmt.Sprintf("PAIEMENTS ABIDJAN.NET DU %v", time.Now().Format("02-01-2006")),
		},
		Body: body,
	}
	//publish the message to the broker
	log.Println("[PUB] publishing event ", topic)
	if err := pubsub.Publish(topic, msg); err != nil {
		theerror := fmt.Sprintf("%v --from payment_service", err)
		log.Printf("[PUB] failed %s\n", theerror)
	}
	return nil
}

func init() {
	/*/*if os.Getenv("ENV") != "PROD" || os.Getenv("ENV") != "TEST" {
		if err := godotenv.Load("../.env"); err != nil {
			log.Fatalf("Couldnt load .env file %v", err)
		}
	}*/*/

}

func main() {
	db, err := createPostgresDBConnection()
	defer db.Close()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	} else {
		log.Println("Connected to DB successfully")
	}
	// Automatically migrates the product struct
	// into database columns/types etc. This will
	// check for changes and migrate them each time
	// this service is restarted.
	db.AutoMigrate(&pb.Payment{})
	repo := newPayRepository(db)
	service := micro.NewService(micro.Name("abjnet.service.payment"), micro.WrapHandler(AuthWrapper))
	service.Init()
	//get the broker instance
	pubsub := broker.NewBroker()
	if err := pubsub.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := pubsub.Connect(); err != nil {
		log.Fatal(err)
	}
	//when we receive the taskrunner.deletesubs event we get all payments from DB
	//and we send it to the email service if properly sent we then clear the db
	_, err = pubsub.Subscribe(topic, func(p broker.Event) error {
		//getting all subscription from database
		log.Println("[SUB] receiving event ", topic)
		subs, err := repo.GetAll()
		if err != nil {
			theerror := fmt.Sprintf("%v --from payment_service", err)
			return errors.New(theerror)
		}
		//publishing the event and sending all the subs to the email_service
		if err := publishEvent(subs, pubsub, "payment.sendmail"); err != nil {
			return err
		}
		//now deleting all the subs from the DB
		if _, err := repo.DeleteAll(); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Println(err)
	}

	pb.RegisterPaymentServiceHandler(service.Server(), newPaymentService(repo))
	if err := service.Run(); err != nil {
		theerror := fmt.Sprintf("%v --from payment_service", err)
		fmt.Println(theerror)
	}
}
