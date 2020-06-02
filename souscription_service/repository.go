package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	pb "github.com/zjjt/abjnet/souscription_service/proto/souscription"
)

type repository interface {
	Subscribe(sub *pb.Souscription) (*pb.Souscription, error)
	GetAll(etat string) ([]*pb.Souscription, error)
	DeleteAll(etat string) (bool, error)
	UpdateAll(etatactuel string, newetat string) (bool, error)
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

//UpdateAll updates all subs where etat = args
func (repo *SubRepository) UpdateAll(etatactuel string, newetat string) (bool, error) {
	if err := repo.db.Model(pb.Souscription{}).Where("etattraitement = ?", etatactuel).Updates(pb.Souscription{Etattraitement: newetat}).Error; err != nil {
		return false, err
	}
	return true, nil
}

//GetAll gets all the subscription in db based on the etattraitement if it is set
func (repo *SubRepository) GetAll(etat string) ([]*pb.Souscription, error) {
	var subs []*pb.Souscription
	if etat == "CREE" || etat == "TRAITEMENT" {
		if err := repo.db.Find(&subs, "etattraitement = ?", etat).Error; err != nil {
			return nil, err
		}
	} else {
		if err := repo.db.Find(&subs).Error; err != nil {
			return nil, err
		}
	}

	return subs, nil

}

//DeleteAll deletes all the subscriptions in db based on the etattraitement=TRAITEE if it is set or just removes everything
func (repo *SubRepository) DeleteAll(etat string) (bool, error) {
	log.Println("deleting subs now")
	err := repo.db.Exec("delete from souscriptions where etattraitement = ?", etat).Error
	log.Println(err)
	if err != nil {
		return false, err
	}

	return true, nil

}
