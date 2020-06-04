package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	pb "github.com/zjjt/abjnet/payment_service/proto/payment"
)

type repository interface {
	MakePayment(payment *pb.Payment) (*pb.Payment, error)
	GetAll() ([]*pb.Payment, error)
	DeleteAll() (bool, error)
}

//PayRepository implements the repository interface
type PayRepository struct {
	db *gorm.DB
}

func newPayRepository(db *gorm.DB) *PayRepository {
	return &PayRepository{db}
}

//MakePayment creates a payment in the DB
func (repo *PayRepository) MakePayment(payment *pb.Payment) (*pb.Payment, error) {
	payTime := time.Now().Format("02-01-2006 15:04")
	payment.CreatedAt = payTime
	//check if the payment doesnt exist already
	payexist := new(pb.Payment)
	log.Println("payexist is:", payexist)
	if err := repo.db.FirstOrCreate(&payexist, payment).Error; err != nil {
		log.Println(fmt.Errorf("payexist:%v", payexist))
		return nil, fmt.Errorf("erreur: %v", err)
	}

	return payexist, nil
}

//GetAll gets all the payments in db
func (repo *PayRepository) GetAll() ([]*pb.Payment, error) {
	var pays []*pb.Payment
	if err := repo.db.Find(&pays).Error; err != nil {
		return nil, err
	}

	return pays, nil

}

//DeleteAll deletes all the payments
func (repo *PayRepository) DeleteAll() (bool, error) {
	log.Println("deleting payments now")
	err := repo.db.Exec("delete from prestations").Error
	log.Println(err)
	if err != nil {
		return false, err
	}

	return true, nil

}
