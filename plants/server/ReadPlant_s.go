package main

import (
	"context"
	pb "github.com/mbcarruthers/gRPCplants/plants/proto"
	"log"
)

// ReadPlant reads a single plant in the database by finding it with the generated UUID
// from the database.
func (s *PlantService) ReadPlant(ctx context.Context, in *pb.PlantId) (*pb.Plant, error) {
	log.Println("ReadPlant(server) invoked")
	queryPlant := Plant{
		Id: in.Id,
	}
	err := s.session.Collection("native.plants").Find().One(&queryPlant)
	if err != nil {
		return nil, err
	}

	return plantToProto(queryPlant), nil
}
