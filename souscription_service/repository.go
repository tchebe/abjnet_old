package main

import (
	"github.com/jinzhu/gorm"
	pb "github.com/zjjt/abjnet/souscription_service/proto/souscription"
)

type repository interface {
	Subscribe(sub *pb.Souscription) error
	GetAll() ([]*pb.Souscription, error)
	DeleteAll() (bool, error)
}

type SubRepository struct {
	db *gorm.DB
}

func newSubRepository(db *gorm.DB) *SubRepository {
	return &SubRepository{db}
}
func (repo *SubRepository) Subscribe(sub *pb.Souscription) error {
	if err := repo.db.Create(sub).Error; err != nil {
		return err
	}
	return nil
}

func (repo *SubRepository) GetAll() ([]*pb.Souscription, error) {
	var subs []*pb.Souscription
	if err := repo.db.Find(&subs).Error; err != nil {
		return nil, err
	}
	return subs, nil

}
func (repo *SubRepository) DeleteAll() (bool, error) {
	if err := repo.db.Exec("TRUNCATE TABLE souscriptions RESTART IDENTITY;").Error; err != nil {
		return false, err
	}
	return true, nil

}
