package main

import (
	"github.com/jinzhu/gorm"
	pb "github.com/zjjt/abjnet/product_service/proto/product"
)

type repository interface {
	Create(product *pb.Product) error
	Get(id string) (*pb.Product, error)
	GetAll() ([]*pb.Product, error)
	Delete(id string) (bool, error)
}

type ProductRepository struct {
	db *gorm.DB
}

func newProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}
func (repo *ProductRepository) Create(product *pb.Product) error {
	if err := repo.db.Create(product).Error; err != nil {
		return err
	}
	return nil
}
func (repo *ProductRepository) Get(id string) (*pb.Product, error) {
	var product *pb.Product
	product.Id = id
	if err := repo.db.First(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}
func (repo *ProductRepository) GetAll() ([]*pb.Product, error) {
	var products []*pb.Product
	if err := repo.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil

}
