package main

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/zjjt/abjnet/souscription_service/proto/souscription"
)

type service struct {
	repo repository
}

func newSouscriptionService(repo repository) *service {
	return &service{repo}
}

//Subscribe -returns the souscription inserted in the DB
func (s *service) Subscribe(ctx context.Context, req *pb.Souscription, res *pb.Response) error {
	if err := s.repo.Subscribe(req); err != nil {
		theerror := fmt.Sprintf("%v --from souscription_service", err)
		res.Done = false
		res.Errors = []*pb.Error{{
			Code:        400,
			Description: "souscription echouée.Veuillez contacter l'administrateur système",
		}}
		return errors.New(theerror)
	}
	//get the id of the subscription
	resp, err := s.repo.GetSub(req)
	if err != nil {
		theerror := fmt.Sprintf("%v --from souscription_service", err)
		res.Done = false
		res.Errors = []*pb.Error{{
			Code:        400,
			Description: "souscription echouée.Veuillez contacter l'administrateur système",
		}}
		return errors.New(theerror)
	}
	res.Done = true
	res.Description = "Souscription prise en compte.Un retour vous sera fait d'ici 24h"
	res.Souscription.Id = resp.Id
	return nil
}

//GetAll -returns a slice of souscriptions
func (s *service) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	subs, err := s.repo.GetAll("")
	if err != nil {
		theerror := fmt.Sprintf("%v --from souscription_service", err)
		return errors.New(theerror)
	}
	res.Souscriptions = subs
	return nil
}

//DeleteAll -deletes all subs from the database
func (s *service) DeleteAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	done, err := s.repo.DeleteAll("")
	if err != nil {
		theerror := fmt.Sprintf("%v --from souscription_service", err)
		return errors.New(theerror)
	}
	res.Done = done
	return nil
}
