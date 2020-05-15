package main

import (
	"os"

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
	if os.Getenv("IN_NSIA") == "no" {
		return &pb.Product{Id: "1", Name: "CAREC TEST RETRAITE"}, nil
	}
	product.Id = id
	if err := repo.db.First(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}
func (repo *ProductRepository) GetAll() ([]*pb.Product, error) {
	var products []*pb.Product
	if os.Getenv("IN_NSIA") == "no" {
		products = append(products, &pb.Product{Id: "1", Name: "CAREC TEST RETRAITE"})
		products = append(products, &pb.Product{Id: "2", Name: "CAREC TEST EPARGNE"})
	} else {
		req := repo.db.Raw("exec dbo.lstprdweblogy").Scan(&products)
		if err := req.Error; err != nil {
			return nil, err
		}

	}
	return products, nil
}
