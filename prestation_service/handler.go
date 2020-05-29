package main

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/zjjt/abjnet/prestation_service/proto/prestation"
)

type service struct {
	repo repository
}

func newPrestationService(repo repository) *service {
	return &service{repo}
}

func (s *service) Rachat(ctx context.Context, req *pb.Prestation, res *pb.Response) error {
	resp, err := s.repo.MakeRachat(req)
	if err != nil {
		theerror := fmt.Sprintf("%v --from prestation_service", err)
		res.Done = false
		res.Errors = []*pb.Error{{
			Code:        400,
			Description: "prestation echouée.Veuillez contacter l'administrateur système",
		}}
		return errors.New(theerror)
	}

	res.Done = true
	res.Description = "Prestation prise en compte.Un retour vous sera fait d'ici 24h"
	res.Prestation = resp
	return nil
}
func (s *service) ValeurRachat(ctx context.Context, req *pb.Request, res *pb.Response) error {
	res.Done = true
	res.Maximumrachetable = "100000"
	return nil
}
