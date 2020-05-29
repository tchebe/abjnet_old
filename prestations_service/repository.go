package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	pb "github.com/zjjt/abjnet/prestation_service/proto/prestation"
)

type repository interface {
	MakePayment(prestation *pb.Prestation) (*pb.Prestation, error)
	GetAll() ([]*pb.Prestation, error)
	DeleteAll() (bool, error)
}
type PrestaRepository struct {
	db *gorm.DB
}

func newPrestaRepository(db *gorm.DB) *PrestaRepository {
	return &PrestaRepository{db}
}

//MakePayment creates a payment in the DB
func (repo *PrestaRepository) MakePrestation(presta *pb.Prestation) (*pb.Prestation, error) {
	prestaTime := time.Now().Format("02-01-2006 15:04")
	presta.CreatedAt = prestaTime
	//check if the Prestation doesnt exist already
	prestaexist := new(pb.Prestation)
	if err := repo.db.FirstOrCreate(&prestaexist, presta).Error; err != nil {
		fmt.Printf("payexist:%v", prestaexist)
		return nil, fmt.Errorf("Cette prestation existe déjà")
	}

	return prestaexist, nil
}
func (repo *PrestaRepository) checkElligibility(presta *pb.Prestation) error {
	//check the remaining montant from the last row in db
	p := new(pb.Prestation)
	if err := repo.db.Last(&p, "nomclient = ? and prenomclient = ? and policeno = ?",presta.).Error; err != nil {
		log.Println(err)
	}

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
