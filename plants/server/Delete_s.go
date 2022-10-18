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

func (s *PlantService) DeletePlant(ctx context.Context, plantId *pb.PlantId) (*empty.Empty, error) {
	log.Println("DeletePlant(server) invoked")

	if err := s.session.Collection("native.plants").Find(db.Cond{"id": plantId.Id}).Delete(); err != nil {
		return &empty.Empty{}, status.Errorf(codes.NotFound, // return code [5]
			fmt.Sprintf("Error finding data [%d]\n %s\n",
				codes.NotFound,
				err.Error()))
	} else {
		return &empty.Empty{}, nil
	}
}
