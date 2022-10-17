package main

import (
	"context"
	"fmt"
	pb "github.com/mbcarruthers/gRPCplants/plants/proto"
	"log"
)

// CreatePlants: Inserts a new Plant in the cockroach database, disregards any pre-implementation of an id
// An Id is created that is IETF REF-4122 complient.
func (s *PlantService) CreatePlant(ctx context.Context, plant *pb.Plant) (*pb.PlantId, error) {
	log.Println("CreatePlant(server) invoked")

	insertPlant := &Plant{
		CommonName: plant.CommonName,
		Genus:      plant.Genus,
		Species:    plant.Species,
	}

	res, err := s.session.Collection("native.plants").Insert(insertPlant)

	if err != nil {
		log.Printf("Error inserting %+v \n %s\n",
			insertPlant,
			err.Error())
		return nil, err
	} else {
		return &pb.PlantId{
			Id: fmt.Sprintf("%v", res.ID()),
		}, nil
	}
}
