package main

import (
	"context"
	pb "github.com/mbcarruthers/gRPCplants/plants/proto"
	"log"
)

// updatePlant calls the server-side updatePlant that updates the given plant
// based upon its Id. Returns the updated plant if successful. If not successful
// it will return (nil, error).
func updatePlant(c pb.PlantServiceClient, plant *pb.Plant) (*pb.Plant, error) {
	log.Println("updatePlant(client) called")
	_, err := c.UpdatePlant(context.Background(), plant)
	if err != nil {
		log.Printf("Error updating \n %s\n",
			err.Error())
		return nil, err
	} else {
		return plant, nil
	}
}
