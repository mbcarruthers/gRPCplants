package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/mbcarruthers/gRPCplants/plants/proto"
	"github.com/upper/db/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

// UpdatePlant updates a Plant in the 'plants' table within the 'native' database.
// Looks the plant up by the it's Id(UUID Primary Key) in the 'plants' table and returns
// (&empty.Empty{},nil) if successful and (nil,err) if not.
// if err code == 5 the item could not be found.
func (s *PlantService) UpdatePlant(ctx context.Context, plant *pb.Plant) (*empty.Empty, error) {
	log.Println("UpdatePlant(server) invoked")

	if err := s.session.Collection("native.plants").Find(db.Cond{"id": plant.Id}).Update(protoToPlant(plant)); err != nil {
		return nil, status.Errorf(codes.NotFound, // return code [5]
			fmt.Sprintf("Error updating [%s]\n %s\n",
				plant.Id,
				err.Error()))
	}
	return &empty.Empty{}, nil
}
