package main

import (
	"context"
	pb "github.com/mbcarruthers/gRPCplants/plants/proto"
	"log"
)

// Client-side inplementation to insert a plant, does not
// Does not need an id, gen_random_uuid() will create that
// in the CockroachDB server.
func createPlant(c pb.PlantServiceClient, plant *pb.Plant) (*pb.PlantId, error) {
	log.Println("createPlant(client) invoked")

	plantId, err := c.CreatePlant(context.Background(), plant)
	if err != nil {
		return nil, err
	}
	return plantId, nil
}
