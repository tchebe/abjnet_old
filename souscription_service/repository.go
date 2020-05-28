package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	pb "github.com/zjjt/abjnet/souscription_service/proto/souscription"
)

type repository interface {
	Subscribe(sub *pb.Souscription) (*pb.Souscription, error)
	GetAll(string) ([]*pb.Souscription, error)
	DeleteAll(string) (bool, error)
	GetSub(sub *pb.Souscription) (*pb.Souscription, error)
}

type SubRepository struct {
	db *gorm.DB
}

func newSubRepository(db *gorm.DB) *SubRepository {
	return &SubRepository{db}
}

//GetSub gets a subscription
func (repo *SubRepository) GetSub(sub *pb.Souscription) (*pb.Souscription, error) {
	sup := new(pb.Souscription)
	if err := repo.db.First(&sup, "nom = ? and prenom = ? and dateofbirth = ? and telephone = ? and abjcardno = ? and montant = ? and codeproduit = ? and datepayment = ? and echeance = ? and beneficiaire = ? and email = ? and etattraitement = ? and created_at = ? ", sub.Nom, sub.Prenom, sub.Dateofbirth, sub.Telephone, sub.Abjcardno, sub.Montant, sub.Codeproduit, sub.Datepayment, sub.Echeance, sub.Beneficiaire, sub.Email, sub.Etattraitement, sub.CreatedAt).Error; err != nil {
		return nil, err
	}
	return sup, nil
}

//Subscribe creates a subscription in the DB
func (repo *SubRepository) Subscribe(sub *pb.Souscription) (*pb.Souscription, error) {
	subTime := time.Now().Format("02-01-2006 15:04")
	sub.CreatedAt = subTime
	//check if the subscription doesnt exist already
	subexist := new(pb.Souscription)
	if err := repo.db.FirstOrCreate(&subexist, sub).Error; err != nil {
		fmt.Printf("subexist:%v", subexist)
		return nil, fmt.Errorf("Cette souscription existe déjà")
	}

	return subexist, nil
}

//GetAll gets all the subscription in db based on the etattraitement if it is set
func (repo *SubRepository) GetAll(etat string) ([]*pb.Souscription, error) {
	var subs []*pb.Souscription
	if etat != "CREE" || etat != "TRAITEMENT" || etat != "TRAITEE" {
		if err := repo.db.Find(&subs).Error; err != nil {
			return nil, err
		}
	} else {
		sub := &pb.Souscription{Etattraitement: etat}
		if err := repo.db.Find(&subs, sub).Error; err != nil {
			return nil, err
		}
	}

	return subs, nil

}

//DeleteAll deletes all the subscriptions in db based on the etattraitement=TRAITEE if it is set or just removes everything
func (repo *SubRepository) DeleteAll(etat string) (bool, error) {
	if etat != "TRAITEE" {
		if err := repo.db.Exec("TRUNCATE TABLE souscriptions RESTART IDENTITY;").Error; err != nil {
			return false, err
		}
	} else {
		if err := repo.db.Where("etattraitement LIKE ?", fmt.Sprintf("%%v%", etat)).Delete(pb.Souscription{}).Error; err != nil {
			return false, err
		}
	}

	return true, nil

}
