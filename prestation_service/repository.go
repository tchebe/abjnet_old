package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	pb "github.com/zjjt/abjnet/prestation_service/proto/prestation"
)

type repository interface {
	MakeRachat(prestation *pb.Prestation) (*pb.Prestation, error)
	GetVR(presta *pb.Prestation) (string, error)
	GetAll(etat string) ([]*pb.Prestation, error)
	UpdateAll(etat string, newetat string) (bool, error)
	DeleteAll(etat string) (bool, error)
}
type PrestaRepository struct {
	db *gorm.DB
}

func newPrestaRepository(db *gorm.DB) *PrestaRepository {
	return &PrestaRepository{db}
}

//MakeRachat creates a demande de prestation in the DB
func (repo *PrestaRepository) MakeRachat(presta *pb.Prestation) (*pb.Prestation, error) {
	prestaTime := time.Now().Format("02-01-2006 15:04")
	presta.CreatedAt = prestaTime

	prestaexist := new(pb.Prestation)
	//check if la demande es elligible ou pas
	if err := repo.checkElligibility(presta); err != nil {
		log.Println("checkelligibility result: ", err)
		return nil, err
	}
	//check if the Prestation doesnt exist
	if err := repo.db.FirstOrCreate(&prestaexist, presta).Error; err != nil {
		fmt.Printf("payexist:%v", prestaexist)
		return nil, fmt.Errorf("Cette prestation existe déjà")
	}

	return prestaexist, nil
}

//GetVR get the VR of a police
func (repo *PrestaRepository) GetVR(presta *pb.Prestation) (string, error) {
	//ici on fait un appel au ws de valeur de rachat avec la police
	//pour avoir le montant rachetable dans le systeme
	//TODO
	//si le ws marche pas on verifie la derniere prestation effectuée avec l'etat TRAITEE
	p := new(pb.Prestation)
	log.Println("police for GetVR in repository is ", presta)
	if presta.Nomclient == "" {
		err := repo.db.Last(&p, "policeno = ? and etattraitement = ?", presta.Policeno, "TRAITEE").Error
		if err != nil {
			//comme le ws de valeur de rachat n'est pas encore disponible on va renvoyer un
			//montant de un million par defaut pour les besoins du test
			//sinon on va renvoyer "",err
			return "1000000", nil
		}
		return p.Montantrestant, nil

	}
	if err := repo.db.Last(&p, "nomclient = ? and prenomclient = ? and policeno = ? and etattraitement = ?", presta.Nomclient, presta.Prenomclient, presta.Policeno, "TRAITEE").Error; err != nil {
		log.Println("in GetVR ", err)
		//comme le ws de valeur de rachat n'est pas encore disponible on va renvoyer un
		//montant de un million par defaut pour les besoins du test
		//sinon on va renvoyer "",err
		return "1000000", nil
	}

	return p.Montantrestant, nil
}
func (repo *PrestaRepository) checkElligibility(presta *pb.Prestation) error {
	//check the remaining montant from the last row in db which hasnt been treated yet
	p := new(pb.Prestation)
	err := repo.db.Last(&p, "nomclient = ? and prenomclient = ? and policeno = ? and etattraitement = ? or etattraitement = ?", presta.Nomclient, presta.Prenomclient, presta.Policeno, "CREE", "TRAITEMENT").Error
	log.Println("in checkelligibility ", err)

	if err == nil {
		//si on a retrouvé une prestation a l'etat CREE on renvoi une erreur
		return errors.New("Une demande de prestation est en cours de traitement,veuillez réessayer plus tard")
	}

	montantD, err := strconv.Atoi(presta.Montantdemande)
	if err != nil {
		return errors.New("erreur pendant la conversion du montant demandé")
	}
	vr, err := repo.GetVR(presta)
	if err != nil {
		return errors.New("erreur pendant l'obtention de la valeur de rachat")
	}
	montantR, err := strconv.Atoi(vr)
	if montantD > montantR {
		return errors.New("erreur le montant demandé est au delà du maximum rachetable possible")
	}

	return nil

}

//GetAll gets all the prestations in db where etat = ?
func (repo *PrestaRepository) GetAll(etat string) ([]*pb.Prestation, error) {
	var p []*pb.Prestation
	if etat != "" {
		if err := repo.db.Find(&p, "etattraitement = ?", etat).Error; err != nil {
			return nil, err
		}
	} else {
		if err := repo.db.Find(&p).Error; err != nil {
			return nil, err
		}
	}

	return p, nil

}

//DeleteAll deletes all the payments where etat = ?
func (repo *PrestaRepository) DeleteAll(etat string) (bool, error) {
	if etat == "" {
		if err := repo.db.Exec("TRUNCATE TABLE prestations RESTART IDENTITY;").Error; err != nil {
			return false, err
		}
	} else {
		if err := repo.db.Where("etattraitement LIKE ?", fmt.Sprintf("%%v%", etat)).Delete(pb.Prestation{}).Error; err != nil {
			return false, err
		}
	}

	return true, nil
}

//UpdateAll updates all subs where etat = args
func (repo *PrestaRepository) UpdateAll(etatactuel string, newetat string) (bool, error) {
	if err := repo.db.Model(pb.Prestation{}).Where("etattraitement = ?", etatactuel).Updates(pb.Prestation{Etattraitement: newetat}).Error; err != nil {
		return false, err
	}
	return true, nil
}
