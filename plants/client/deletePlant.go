package main

import (
	"context"
	pb "github.com/mbcarruthers/gRPCplants/plants/proto"
	"log"
)

func deletePlant(c pb.PlantServiceClient, plantId *pb.PlantId) error {
	if _, err := c.DeletePlant(context.Background(), plantId); err != nil {
		log.Printf("Err on the server\n %s\n",
			err.Error())
		return err
	} else {
		return nil
	}
}
