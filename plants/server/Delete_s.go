package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/mbcarruthers/gRPCplants/plants/proto"
	"github.com/upper/db/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Deletes a Plant in the database by finding it with the plants id. If there in an error it returns
// (&empty.Empty{}, err) if it was successful it returns (&empty.Empty{}, nil)
// error code 5 means that the plant was not found, either it was deleted already or there is something
// wrong with the id.
func (s *PlantService) DeletePlant(ctx context.Context, plantId *pb.PlantId) (*empty.Empty, error) {

	if err := s.session.Collection("native.plants").Find(db.Cond{"id": plantId.Id}).Delete(); err != nil {
		return &empty.Empty{}, status.Errorf(codes.NotFound, // return code [5]
			fmt.Sprintf("Error finding data [%d]\n %s\n",
				codes.NotFound,
				err.Error()))
	} else {
		return &empty.Empty{}, nil
	}
}
