package main

import (
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/mbcarruthers/gRPCplants/plants/proto"
	"log"
)

// ListPlants lists all plants that are stored in the plants table within the native database.
func (s *PlantService) ListPlants(in *empty.Empty, stream pb.PlantService_ListPlantsServer) error {
	log.Println("List Plants(server) invoked")
	plants := []Plant{}
	err := s.session.Collection("native.plants").Find().All(&plants)

	if err != nil {
		log.Fatalf("Error With collection.Find(): \n %s\n",
			err.Error())
	}
	// stream plants to the client
	for _, plant := range plants {
		if err := stream.Send(plantToProto(plant)); err != nil {
			log.Printf("Error streaming  %+v \n %s\n",
				&plant, err.Error())
		}
	}

	return nil
}
