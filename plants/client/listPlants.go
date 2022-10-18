package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/mbcarruthers/gRPCplants/plants/proto"
	"io"
	"log"
)

// listPlants: Lists the plants within the database thus far.
// As of now only prints the plants to the console
func listPlants(c pb.PlantServiceClient) {
	log.Println("List plants(client) invoked")

	stream, err := c.ListPlants(context.Background(), &empty.Empty{})
	if err != nil {
		log.Fatalf("Error while calling list plants \n %s \n",
			err.Error())
	}
	// stream content from the server
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Something wen't wrong\n %s\n",
				err.Error())
		} else {
			log.Printf("%+v\n",
				res)
		}
	}
}
