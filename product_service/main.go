package main

import (
	"fmt"
	"log"

	"github.com/micro/go-micro/v2"
	pb "github.com/zjjt/abjnet/product_service/proto/product"
)

func main() {
	db, err := createSqlServerDBConnection()
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
	pb.RegisterProductServiceHandler(service.Server(), newProductService(repo))
	if err := service.Run(); err != nil {
		theerror := fmt.Sprintf("%v --from product_service", err)
		fmt.Println(theerror)
	}
}
