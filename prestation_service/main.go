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
	pb "github.com/tchebe/abjnet/prestation_service/proto/prestation"
	userproto "github.com/tchebe/abjnet/user_service/proto/user"
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
			return errors.New("no auth meta-data found in request --from prestation_service l35")
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
			theerror := fmt.Sprintf("%v --from prestation_service l46", err)
			return errors.New(theerror)
		}
		err = fn(ctx, req, res)
		return err
	}
}

var topic = []string{"taskrunner.updatepresta_traitement", "taskrunner.updatepresta_traitee", "taskrunner.deleteprestations"}

func publishEvent(subs []*pb.Prestation, pubsub broker.Broker, topic string) error {
	//when sending an event we have to serialize it to bytes
	//we are sending to our ecosystem the event prestation.sendmail with the details
	//with all today's subscription
	body, err := json.Marshal(subs)
	if err != nil {
		theerror := fmt.Sprintf("%v --from prestation_service", err)
		return errors.New(theerror)
	}

	//create a broker message
	msg := &broker.Message{
		Header: map[string]string{
			"to":    os.Getenv("TO"),
			"cc":    os.Getenv("CC"),
			"objet": fmt.Sprintf("DEMANDES DE PRESTATION ABIDJAN.NET DU %v", time.Now().Format("02-01-2006")),
		},
		Body: body,
	}
	//publish the message to the broker
	log.Println("[PUB] publishing event ", topic)
	if err := pubsub.Publish(topic, msg); err != nil {
		theerror := fmt.Sprintf("%v --from prestation_service", err)
		log.Printf("[PUB] failed %s\n", theerror)
	}
	return nil
}

func init() {
	/*if os.Getenv("ENV") != "PROD" || os.Getenv("ENV") != "TEST" {
		if err := godotenv.Load("../.env"); err != nil {
			log.Fatalf("Couldnt load .env file %v", err)
		}
	}*/

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
	db.AutoMigrate(&pb.Prestation{})
	repo := newPrestaRepository(db)
	service := micro.NewService(micro.Name("abjnet.service.prestation"), micro.WrapHandler(AuthWrapper))
	service.Init()
	//get the broker instance
	pubsub := broker.NewBroker()
	if err := pubsub.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := pubsub.Connect(); err != nil {
		log.Fatal(err)
	}
	//when we receive the taskrunner.updatepresta_traitement event we update all prestations from DB
	//with etattraitement CREE to TRAITEMENT
	//and we send it to the email service
	_, err = pubsub.Subscribe(topic[0], func(p broker.Event) error {
		//updating all prestations from database
		log.Println("[SUB] receiving event ", topic[0])
		_, err := repo.UpdateAll("CREE", "TRAITEMENT")
		if err != nil {
			theerror := fmt.Sprintf("%v --from prestation_service", err)
			return errors.New(theerror)
		}
		//sending all the prestations in TRAITEMENT to NSIA CHAP CHAP BORNE
		//TODO
		//res,err:=http.Get(fmt.Sprintf("%s/saveDemandeExterne?client_id=${dataForm['client_id'].toString()}&historique_connexion_id=${dataForm['historique_connexion_id'].toString()}&numeroPayeur=${dataForm['numeroPayeur'].toString()}&police=${dataForm['police'].toString()}&montant=${dataForm['montant'].toString()}&nbre_mois_remb=${dataForm['nbre_mois_remb'].toString()}&type_demande_id=${dataForm['type_demande_id'].toString()}&provenance=2"))
		//getting all the prestations to send
		prestations, err := repo.GetAll("TRAITEMENT")
		if err != nil {
			theerror := fmt.Sprintf("%v --from prestation_service", err)
			return errors.New(theerror)
		}
		//publishing the event and sending all the subs to the email_service
		if err := publishEvent(prestations, pubsub, "prestation.sendmail"); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Println(err)
	}
	//when we receive the taskrunner.updatepresta_traitee event we update all prestations from DB
	//with etattraitement TRAITEMENT to TRAITEE
	_, err = pubsub.Subscribe(topic[1], func(p broker.Event) error {
		//updating all prestations from database
		log.Println("[SUB] receiving event ", topic[1])
		_, err := repo.UpdateAll("TRAITEMENT", "TRAITEE")
		if err != nil {
			theerror := fmt.Sprintf("%v --from prestation_service", err)
			return errors.New(theerror)
		}
		return nil
	})

	if err != nil {
		log.Println(err)
	}
	//when we receive the taskrunner.deleteprestation event we delete all prestations from DB
	//with etattraitement TRAITEE
	_, err = pubsub.Subscribe(topic[2], func(p broker.Event) error {
		//updating all prestations from database
		log.Println("[SUB] receiving event ", topic[2])
		_, err := repo.DeleteAll("TRAITEE")
		if err != nil {
			theerror := fmt.Sprintf("%v --from prestation_service", err)
			return errors.New(theerror)
		}
		return nil
	})

	if err != nil {
		log.Println(err)
	}

	pb.RegisterPrestationServiceHandler(service.Server(), newPrestationService(repo))
	if err := service.Run(); err != nil {
		theerror := fmt.Sprintf("%v --from payment_service", err)
		fmt.Println(theerror)
	}
}
