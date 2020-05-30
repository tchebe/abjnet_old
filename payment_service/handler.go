package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	pb "github.com/zjjt/abjnet/payment_service/proto/payment"
)

type service struct {
	repo repository
}

func newPaymentService(repo repository) *service {
	return &service{repo}
}

func (s *service) Pay(ctx context.Context, req *pb.Payment, res *pb.Response) error {
	log.Println(req)
	resp, err := s.repo.MakePayment(req)
	if err != nil {
		theerror := fmt.Sprintf("%v --from payment_service", err)
		res.Done = false
		res.Errors = []*pb.Error{{
			Code:        400,
			Description: "paiement echouée.Veuillez contacter l'administrateur système",
		}}
		return errors.New(theerror)
	}

	res.Done = true
	res.Description = "Paiement pris en compte.Un retour vous sera fait d'ici 24h"
	res.Payment = resp
	return nil
}
