package main

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/zjjt/abjnet/product_service/proto/product"
)

//event to be sent
const topic = "product.deleted"

type service struct {
	repo repository
}

func newProductService(repo repository) *service {
	return &service{repo}
}

//Get - retrieves a single product
func (s *service) Get(ctx context.Context, req *pb.Product, res *pb.Response) error {
	product, err := s.repo.Get(req.Id)
	if err != nil {
		theerror := fmt.Sprintf("%v --from product_service", err)
		return errors.New(theerror)
	}
	res.Product = product
	return nil
}

//GetAll -returns a slice of products
func (s *service) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	products, err := s.repo.GetAll()
	if err != nil {
		theerror := fmt.Sprintf("%v --from product_service", err)
		return errors.New(theerror)
	}
	res.Products = products
	return nil
}
