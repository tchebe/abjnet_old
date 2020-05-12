package main

import (
	"github.com/jinzhu/gorm"
	pb "github.com/zjjt/abjnet/product_service/proto/product"
)

type repository interface {
	Get(id string) (*pb.Product, error)
	GetAll() ([]*pb.Product, error)
}

type ProductRepository struct {
	db *gorm.DB
}

func newProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
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
	req := repo.db.Raw("SELECT JAPRODP_WNPRO as CodeProduit, JAPRODP_LIPR01 as Name FROM NSIACIP.JAPRODP").Scan(&products)
	if err := req.Error; err != nil {
		return nil, err
	}
	return products, nil

}
