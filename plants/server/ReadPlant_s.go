package main

import (
	"context"
	pb "github.com/mbcarruthers/gRPCplants/plants/proto"
	"github.com/upper/db/v4"
	"log"
)

// ReadPlant reads a single plant in the database by finding it with the generated UUID
// from the database. Returns the &pb.Plant if found. else it returns an error
func (s *PlantService) ReadPlant(ctx context.Context, in *pb.PlantId) (*pb.Plant, error) {
	log.Println("ReadPlant(server) invoked")
	queryPlant := Plant{
		Id: in.Id,
	}

	err := s.session.Collection("native.plants").Find(db.Cond{"id": queryPlant.Id}).One(&queryPlant)
	if err != nil {
		return nil, err
	}
	// return protobuf
	return plantToProto(queryPlant), nil
}
