package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/robfig/cron/v3"
)

var topic = []string{"taskrunner.deletesubs", "taskrunner.updatesubs", "taskrunner.deletepayments", "taskrunner.updatepresta_traitement", "taskrunner.updatepresta_traitee", "taskrunner.deleteprestations"}

func publishEvent(pubsub broker.Broker, topic string) error {
	//create a broker message
	msg := &broker.Message{}
	//publish the message to the broker
	log.Println("[PUB] publishing event ", topic)
	if err := pubsub.Publish(topic, msg); err != nil {
		theerror := fmt.Sprintf("%v --from taskrunner_service", err)
		log.Printf("[PUB] failed %s\n", theerror)
	}
	return nil
}
func init() {
	if os.Getenv("ENV") != "PROD" || os.Getenv("ENV") != "TEST" {
		if err := godotenv.Load("../.env"); err != nil {
			log.Fatalf("Couldnt load .env file %v", err)
		}
	}

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

	job := cron.New()
	job.AddFunc(os.Getenv("MAJSUBAT"), func() {
		//tells souscription_service to update all the subs to TRAITEMENT and send all them to the email_service

		if err := publishEvent(pubsub, topic[1]); err != nil {
			fmt.Println(err)
		}
		//une fois la mise a jour a l'etat de TRAITEMENT est faite
		//on attend 24h pour vider la tabe des souscriptions
		if os.Getenv("ENV") == "TEST" {
			//code de test
			t, _ := strconv.Atoi(os.Getenv("NBHOURTODELSUB"))
			time.AfterFunc(time.Duration(t)*time.Minute, func() {
				if err := publishEvent(pubsub, topic[0]); err != nil {
					fmt.Println(err)
				}
			})
		} else {
			//code de production
			t, _ := strconv.Atoi(os.Getenv("NBHOURTODELSUB"))
			time.AfterFunc(time.Duration(t)*time.Hour, func() {
				if err := publishEvent(pubsub, topic[0]); err != nil {
					fmt.Println(err)
				}
			})
		}

	})

	job.AddFunc(os.Getenv("DELETEPAYSAT"), func() {
		// la table des paiement est vidées tous les jours a l'heure en parametre
		if err := publishEvent(pubsub, topic[2]); err != nil {
			fmt.Println(err)
		}
	})
	job.AddFunc(os.Getenv("MAJPREAT"), func() {
		//tells prestation_service to update all the prestations to TRAITEMENT and
		//send all them to the email_service

		if err := publishEvent(pubsub, topic[3]); err != nil {
			fmt.Println(err)
		}
		//une fois la mise a jour a l'etat de TRAITEMENT est faite
		//on attend 5jours pour mettre a jour a l'etat TRAITEE
		if os.Getenv("ENV") == "TEST" {
			//code de test
			t, _ := strconv.Atoi(os.Getenv("NBDAYTOUPPRE"))
			time.AfterFunc(time.Duration(t)*time.Minute, func() {
				if err := publishEvent(pubsub, topic[4]); err != nil {
					fmt.Println(err)
				}
				//et apres un certains nombre de jour on vide les prestations TRAITEE de la base
				t, _ := strconv.Atoi(os.Getenv("NBDAYTODELPRE"))
				time.AfterFunc(time.Duration(t)*time.Minute, func() {
					if err := publishEvent(pubsub, topic[5]); err != nil {
						fmt.Println(err)
					}
				})

			})

		} else {
			//code de production
			t, _ := strconv.Atoi(os.Getenv("NBDAYTOUPPRE"))
			time.AfterFunc(time.Duration(t)*24*time.Hour, func() {
				if err := publishEvent(pubsub, topic[4]); err != nil {
					fmt.Println(err)
				}
				//et apres un certains nombre de jour on vide les prestations TRAITEE de la base
				t, _ := strconv.Atoi(os.Getenv("NBDAYTODELPRE"))
				time.AfterFunc(time.Duration(t)*24*time.Hour, func() {
					if err := publishEvent(pubsub, topic[5]); err != nil {
						fmt.Println(err)
					}
				})
			})

		}

	})

	job.Start()
	if err := service.Run(); err != nil {
		theerror := fmt.Sprintf("%v --from taskrunner_service", err)
		fmt.Println(theerror)
	}
	//to stop the goroutine from exiting unless we receive interrupt from
	c := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		log.Println("exiting now...")
		<-c
		os.Exit(0)
	}()
	for {
	}

}
