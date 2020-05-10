package main

import (
	"fmt"
	"log"

	"github.com/micro/go-micro/v2"
	pb "github.com/zjjt/abjnet/product_service/proto/product"
)

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
	db.AutoMigrate(&pb.Product{})
	repo := newProductRepository(db)
	service := micro.NewService(micro.Name("abjnet.service.product"))
	service.Init()
	//get an instance of the event broker
	pubsub := service.Server().Options().Broker
	if err := pubsub.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := pubsub.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}
	pb.RegisterProductServiceHandler(service.Server(), newProductService(repo, pubsub))
	if err := service.Run(); err != nil {
		theerror := fmt.Sprintf("%v --from ProductService", err)
		fmt.Println(theerror)
	}
}
