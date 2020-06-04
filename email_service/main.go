package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/go-mail/mail"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	pbPay "github.com/zjjt/abjnet/payment_service/proto/payment"
	pbPre "github.com/zjjt/abjnet/prestation_service/proto/prestation"
	pbS "github.com/zjjt/abjnet/souscription_service/proto/souscription"
)

func main() {
	//slice of topics to suscribe to
	topics := []string{"user.created", "souscription.sendmail", "payment.sendmail", "prestation.sendmail"}
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

	_, err := pubsub.Subscribe(topics[0], func(p broker.Event) error {
		log.Println("[SUB] receiving event ", topics[0])
		eventHeadersMap := p.Message().Header
		log.Println("event msg ", eventHeadersMap)
		go sendEmail(os.Getenv("FROM"), eventHeadersMap["to"], eventHeadersMap["cc"], eventHeadersMap["objet"], "Bonjour,<br/> un test", p.Message().Body)
		return nil
	})
	log.Println("[SUB ERROR]", err)
	_, err = pubsub.Subscribe(topics[1], func(p broker.Event) error {
		log.Println("[SUB] receiving event ", topics[1])
		eventHeadersMap := p.Message().Header
		log.Println("event msg ", eventHeadersMap)

		go sendEmail(os.Getenv("FROM"), eventHeadersMap["to"], eventHeadersMap["cc"], eventHeadersMap["objet"], "Bonjour,<br/> un test", p.Message().Body)
		return nil
	})
	log.Println("[SUB ERROR]", err)
	_, err = pubsub.Subscribe(topics[2], func(p broker.Event) error {
		log.Println("[SUB] receiving event ", topics[2])
		eventHeadersMap := p.Message().Header
		log.Println("event msg ", eventHeadersMap)

		go sendEmail(os.Getenv("FROM"), eventHeadersMap["to"], eventHeadersMap["cc"], eventHeadersMap["objet"], "Bonjour,<br/> un test", p.Message().Body)
		return nil
	})
	log.Println("[SUB ERROR]", err)
	_, err = pubsub.Subscribe(topics[3], func(p broker.Event) error {
		log.Println("[SUB] receiving event ", topics[3])
		eventHeadersMap := p.Message().Header
		log.Println("event msg ", eventHeadersMap)

		go sendEmail(os.Getenv("FROM"), eventHeadersMap["to"], eventHeadersMap["cc"], eventHeadersMap["objet"], "Bonjour,<br/> un test", p.Message().Body)
		return nil
	})
	log.Println("[SUB ERROR]", err)

	//run the server
	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}
func prepareExcelFileSub(subs []*pbS.Souscription) *excelize.File {
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

		excelfile.SetCellValue("Sheet1", fmt.Sprintf("A%d", i+2), v.Nom)
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("B%d", i+2), v.Prenom)
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("C%d", i+2), v.Dateofbirth)
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("D%d", i+2), v.Telephone)
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("E%d", i+2), v.Abjcardno)
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("F%d", i+2), v.Montant)
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("G%d", i+2), v.Codeproduit)
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("H%d", i+2), v.Datepayment)
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("I%d", i+2), v.Echeance)
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("J%d", i+2), v.Beneficiaire)
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("K%d", i+2), v.Email)

	}
	return excelfile
}

func prepareExcelFilePay(payments []*pbPay.Payment) *excelize.File {
	excelfile := excelize.NewFile()
	//here we create the top header rows
	excelfile.SetCellValue("Sheet1", "A1", "NUMERO TRANSACTION")
	excelfile.SetCellValue("Sheet1", "B1", "NOM ASSURE")
	excelfile.SetCellValue("Sheet1", "C1", "PRENOMS ASSURE")
	excelfile.SetCellValue("Sheet1", "D1", "CONTACT TELEPHONIQUE")
	excelfile.SetCellValue("Sheet1", "E1", "DATE DE PAIEMENT APM")
	excelfile.SetCellValue("Sheet1", "F1", "CONVENTION")
	excelfile.SetCellValue("Sheet1", "G1", "POLICE")
	excelfile.SetCellValue("Sheet1", "H1", "MONTANT")
	excelfile.SetCellValue("Sheet1", "I1", "DATE RECEPTION PAIEMENT NSIA")
	//here we fill the file with the data
	for i, v := range payments {
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("A%d", i+2), v.Transacno)
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("B%d", i+2), v.Nomclient)
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("C%d", i+2), v.Prenomclient)
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("D%d", i+2), v.Telephone)
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("E%d", i+2), v.Datepaymentuser)
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("F%d", i+2), v.Conventionno)
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("G%d", i+2), v.Policeno)
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("H%d", i+2), v.Montant)
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("I%d", i+2), v.CreatedAt)

	}
	return excelfile
}
func prepareExcelFilePresta(prestations []*pbPre.Prestation) *excelize.File {
	excelfile := excelize.NewFile()
	//here we create the top header rows
	excelfile.SetCellValue("Sheet1", "A1", "NUMERO TRANSACTION")
	excelfile.SetCellValue("Sheet1", "B1", "NOM ASSURE")
	excelfile.SetCellValue("Sheet1", "C1", "PRENOMS ASSURE")
	excelfile.SetCellValue("Sheet1", "D1", "CONTACT TELEPHONIQUE")
	excelfile.SetCellValue("Sheet1", "E1", "DATE DE DEMANDE")
	excelfile.SetCellValue("Sheet1", "F1", "CONVENTION")
	excelfile.SetCellValue("Sheet1", "G1", "POLICE")
	excelfile.SetCellValue("Sheet1", "H1", "MONTANT DEMANDE")
	excelfile.SetCellValue("Sheet1", "I1", "MONTANT RESTANT APRES RACHAT")
	excelfile.SetCellValue("Sheet1", "J1", "DATE RECEPTION DEMANDE NSIA")
	//here we fill the file with the data
	for i, v := range prestations {
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("A%d", i+2), v.Transacno)
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("B%d", i+2), v.Nomclient)
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("C%d", i+2), v.Prenomclient)
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("D%d", i+2), v.Telephone)
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("E%d", i+2), v.Datedemandeuser)
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("F%d", i+2), v.Conventionno)
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("G%d", i+2), v.Policeno)
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("H%d", i+2), v.Montantdemande)
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("I%d", i+2), v.Montantrestant)
		excelfile.SetCellValue("Sheet1", fmt.Sprintf("J%d", i+2), v.CreatedAt)

	}
	return excelfile
}
func sendEmail(from string, to string, cc string, topic string, msghtml string, byteArr []byte) error {
	//first we unpack to slices the to and cc args
	TO := strings.Split(to, ",")
	CC := strings.Split(to, ",")
	log.Printf("sending email from %v to %v en cc %v and subject %v", from, TO, CC, topic)
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
		var pays []*pbPay.Payment
		var prestas []*pbPre.Prestation
		var excelfile *excelize.File
		if strings.Contains(topic, "SOUSCRIPTIONS") {
			json.Unmarshal(byteArr, &subs)
			//here we create the excel file from the subs
			excelfile = prepareExcelFileSub(subs)
		} else if strings.Contains(topic, "PRESTATIONS") {
			json.Unmarshal(byteArr, &prestas)
			//here we create the excel file from the subs
			excelfile = prepareExcelFilePresta(prestas)
		} else {
			json.Unmarshal(byteArr, &pays)
			//here we create the excel file from the subs
			excelfile = prepareExcelFilePay(pays)
		}

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
	d := mail.NewDialer(os.Getenv("SMTP_HOST"), port, os.Getenv("FROM"), os.Getenv("ADPASSWORD"))
	d.SSL = false
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
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
