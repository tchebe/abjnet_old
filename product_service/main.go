package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/server"
	pb "github.com/tchebe/abjnet/product_service/proto/product"
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
			return errors.New("no auth meta-data found in request --from product_service l32")
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
			theerror := fmt.Sprintf("%v --from product_service l43", err)
			return errors.New(theerror)
		}
		err = fn(ctx, req, res)
		return err
	}
}

func main() {
	var db *gorm.DB

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
	//db.AutoMigrate(&pb.Product{})
	repo := newProductRepository(db)
	service := micro.NewService(micro.Name("abjnet.service.product"), micro.WrapHandler(AuthWrapper))
	service.Init()
	pb.RegisterProductServiceHandler(service.Server(), newProductService(repo))
	if err := service.Run(); err != nil {
		theerror := fmt.Sprintf("%v --from product_service", err)
		fmt.Println(theerror)
	}
}
