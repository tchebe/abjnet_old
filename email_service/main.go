package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
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
					log.Println("user created")
					return nil
				})
				return err
			case "souscription.sendmail":
				_, err := pubsub.Subscribe(v, func(p broker.Event) error {
					eventHeadersMap := p.Message().Header
					go sendEmail(os.Getenv("FROM"), eventHeadersMap["to"], eventHeadersMap["cc"], eventHeadersMap["objet"], "Bonjour,<br/> un test", p.Message().Body)
					return nil
				})
				return err
			case "product.deleted":
				_, err := pubsub.Subscribe(v, func(p broker.Event) error {
					var product *pbP.Product
					if err := json.Unmarshal(p.Message().Body, &product); err != nil {
						theerror := fmt.Sprintf("%v --from email_service", err)
						return errors.New(theerror)
					}
					log.Println(product)
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
func prepareExcelFile(subs []*pbS.Souscription) *excelize.File {
	excelfile := excelize.NewFile()
	//here we create the top header rows
	excelfile.SetCellValue("Sheet1", "A1", "NOM ASSURE")
	excelfile.SetCellValue("Sheet1", "B1", "PRENOMS ASSURE")
	excelfile.SetCellValue("Sheet1", "C1", "DATE DE NAISSANCE")
	excelfile.SetCellValue("Sheet1", "D1", "CONTACT TELEPHONIQUE")
	excelfile.SetCellValue("Sheet1", "E1", "N CARTE ABIDJAN.NET")
	excelfile.SetCellValue("Sheet1", "F1", "MONTANT PAIEMENT")
	excelfile.SetCellValue("Sheet1", "G1", "CODE PRODUIT")
	excelfile.SetCellValue("Sheet1", "H1", "DATE PAIEMENT")
	excelfile.SetCellValue("Sheet1", "I1", "ECHEANCE")
	excelfile.SetCellValue("Sheet1", "J1", "NOM BENEFICIAIRE")
	excelfile.SetCellValue("Sheet1", "K1", "EMAIL")
	//here we fill the file with the data
	for i, v := range subs {
		index := i + 2
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("A%d", index), v.Nom)
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("B%d", index), v.Prenom)
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("C%d", index), v.Dateofbirth)
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("D%d", index), v.Telephone)
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("E%d", index), v.Abjcardno)
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("F%d", index), v.Montant)
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("G%d", index), v.Codeproduit)
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("H%d", index), v.Datepayment)
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("I%d", index), v.Echeance)
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("J%d", index), v.Beneficiaire)
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("K%d", index), v.Email)

	}
	return excelfile
}

func sendEmail(from string, to string, cc string, topic string, msghtml string, byteArr []byte) error {
	//first we unpack to slices the to and cc args
	TO := strings.Split(to, ",")
	CC := strings.Split(to, ",")

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
		var subs []*pbS.Souscription
		json.Unmarshal(byteArr, &subs)
		//here we create the excel file from the subs
		excelfile := prepareExcelFile(subs)
		// Save xlsx file by the given path.
		if err := excelfile.SaveAs(fmt.Sprintf("%s.xlsx", topic)); err != nil {
			fmt.Println(err)
		}
		//add the excel file to the mail
		m.Attach(fmt.Sprintf("%s.xlsx", topic))
	}
	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		log.Fatal("Error please check the smtp port in environment")
	}
	d := mail.NewDialer(os.Getenv("SMTP_HOST"), port, os.Getenv("FROM"), "ZjjTEnt@1988")
	//d.StartTLSPolicy = mail.MandatoryStartTLS

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		log.Fatal(err)
	}
	//delete the excel file
	if err := os.Remove(fmt.Sprintf("%s.xlsx", topic)); err != nil {
		log.Println(err)
	}
	return nil
}
