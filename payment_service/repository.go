package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	pb "github.com/zjjt/abjnet/payment_service/proto/payment"
)

type repository interface {
	MakePayment(payment *pb.Payment) (*pb.Payment, error)
	GetAll() ([]*pb.Payment, error)
	DeleteAll() (bool, error)
}
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
	if err := repo.db.FirstOrCreate(&payexist, payment).Error; err != nil {
		fmt.Printf("payexist:%v", payexist)
		return nil, fmt.Errorf("erreur: %v\n%v\n%v", err, payment, payexist)
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
	if err := repo.db.Exec("TRUNCATE TABLE payments RESTART IDENTITY;").Error; err != nil {
		return false, err
	}

	return true, nil

}
