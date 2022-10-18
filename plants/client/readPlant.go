package main

import (
	"context"
	pb "github.com/mbcarruthers/gRPCplants/plants/proto"
	"log"
)

// readPlant client implemenation of the readPlant gRPC function to retrieve
// a single plant from the database given its UUID
func readPlant(c pb.PlantServiceClient, plantId *pb.PlantId) *pb.Plant {
	log.Println("readPlant(client) invoked")
	resultPlant, err := c.ReadPlant(context.Background(), plantId)

	// if there is a server-side error finding the plant
	if err != nil {
		log.Fatalf("Error finding plant\n %s\n",
			err.Error())
	}
	return resultPlant
}
