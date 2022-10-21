package main

import (
	"context"
	"fmt"
	pb "github.com/mbcarruthers/gRPCplants/plants/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreatePlants: Inserts a new Plant in the cockroach database, disregards any pre-implementation of an id
// An Id is created that is IETF REF-4122 complient.
func (s *PlantService) CreatePlant(ctx context.Context, plant *pb.Plant) (*pb.PlantId, error) {
	insertPlant := &Plant{
		CommonName: plant.CommonName,
		Genus:      plant.Genus,
		Species:    plant.Species,
	}

	res, err := s.session.Collection("native.plants").Insert(insertPlant)
	if err != nil {
		return nil, status.Errorf(codes.Internal, // return code [13]
			fmt.Sprintf("Error inserting\n %s\n",
				err.Error()))
	} else {
		return &pb.PlantId{
			Id: fmt.Sprintf("%v", res.ID()),
		}, nil
	}
}
