package main

import (
	"log"
	"os"
	"strconv"

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
	type p struct {
		Id   int
		Name string
	}
	var pro []p
	if os.Getenv("IN_NSIA") == "no" {
		products = append(products, &pb.Product{Id: "1", Name: "CAREC TEST RETRAITE"})
		products = append(products, &pb.Product{Id: "2", Name: "CAREC TEST EPARGNE"})
	} else {
		rows, err := repo.db.Debug().Raw("exec dbo.lstprdweblogy").Rows()
		defer rows.Close()
		if err != nil {
			log.Printf("error %v\n", err)
			return nil, err
		}
		for rows.Next() {
			log.Println(rows)
			repo.db.ScanRows(rows, &pro)
		}
		for _, v := range pro {
			products = append(products, &pb.Product{Id: strconv.Itoa(v.Id), Name: v.Name})
		}
		log.Printf("content of pro %v\n", pro)
		log.Printf("content of products %v\n", products)

	}
	if len(products) > 0 {
		return products, nil
	}
	return nil, nil

}
