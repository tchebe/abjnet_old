package main

import (
	"context"
	"log"
	"os"

	"github.com/micro/go-micro/v2"
	microclient "github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/config/cmd"
	pb "github.com/zjjt/abjnet/user_service/proto/user"
)

type user struct {
	name     string
	email    string
	password string
	company  string
}

func main() {

	cmd.Init()
	client := pb.NewUserService("abjnet.service.user", microclient.DefaultClient)
	//utilisation de la ligne de commande pour passer des parametres via l'outil de go-micro
	service := micro.NewService(micro.Name("abj.clitool.user"))
	//start the service
	service.Init()
	userArray := make([]user,1)
	//ici on prerempli les elements du slice de user afin de boucler sur leur creation
	for i:=0;i<1;i++{
		append(userArray,&user{
			name := "weblogie"
			email := "thibaut.zehi@groupensia.com"
			password := "nsi@weblogie130"
			company := "WEBLOGIE"
		})
	}
	
	//here we call our user service
	if len(userArray){
		r, err := client.Create(context.TODO(), &pb.User{
			Name:     name,
			Email:    email,
			Password: password,
			Company:  company,
		})
		if err != nil {
			log.Fatalf("couldnt create user %v", err)
		}
		log.Printf("Created: %s", r.User.Id)
		getAll, err := client.GetAll(context.Background(), &pb.Request{})
		if err != nil {
			log.Fatalf("Couldnt retrieve list of users %v", err)
		}
		for _, v := range getAll.Users {
			log.Println(v)
		}
		authResponse, err := client.Auth(context.Background(), &pb.User{
			Email:    email,
			Password: password,
		})
	
		if err != nil {
			log.Fatalf("Could not authenticate user: %s error: %v\n", email, err)
		}
		log.Printf("Your access token is: %s \n", authResponse.Token)
		os.Exit(0)
	}else{
		log.Fatalf("Could not create users in array\n")
	}
	

	

}
