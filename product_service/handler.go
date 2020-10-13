package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

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
func (s *service) GetClientProducts(ctx context.Context, req *pb.Client, res *pb.Response) error {
	return nil
}
func (s *service) GetCotisations(ctx context.Context, req *pb.Police, res *pb.Response) error {
	resp, err := http.Get(fmt.Sprintf("http://10.11.100.48:8084/etatCotisation/%s", req.Police))
	if err != nil {
		theerror := fmt.Sprintf("%v --from product_service", err)
		return errors.New(theerror)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		theerror := fmt.Sprintf("%v --from product_service", err)
		return errors.New(theerror)
	}
	res.Etat = string(body)
	return nil
}

//permettant d'identifier le client de weblogy
func (s *service) GetlistePoliceExterne(ctx context.Context, req *pb.Police, res *pb.Response) error {
	resp, err := http.Get(fmt.Sprintf("http://10.11.100.48:8084/listePoliceExterne/%s", req.Police))
	if err != nil {
		theerror := fmt.Sprintf("%v --from product_service", err)
		return errors.New(theerror)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		theerror := fmt.Sprintf("%v --from product_service", err)
		return errors.New(theerror)
	}
	res.Etat = string(body)
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
