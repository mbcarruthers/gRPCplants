package main

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/mbcarruthers/gRPCplants/plants/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

// ListPlants lists all plants that are stored in the plants table within the native database.
// if error code is 2 then there was some sort of error streaming the data to the client.
// if error code is 5 then the data was not found.
func (s *PlantService) ListPlants(in *empty.Empty, stream pb.PlantService_ListPlantsServer) error {
	log.Println("List Plants(server) invoked")
	plants := []Plant{}
	err := s.session.Collection("native.plants").Find().All(&plants)

	if err != nil {
		return status.Errorf(codes.NotFound,
			fmt.Sprintf("Error Not Found\n %s\n",
				err.Error()))
	}
	// stream plants to the client
	for _, plant := range plants {
		if err := stream.Send(plantToProto(plant)); err != nil {
			return status.Errorf(codes.Unknown, // return code [2]
				fmt.Sprintf("Error streaming\n %s\n",
					err.Error()))
		}
	}

	return nil
}
