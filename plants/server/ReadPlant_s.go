package main

import (
	"context"
	"fmt"
	pb "github.com/mbcarruthers/gRPCplants/plants/proto"
	"github.com/upper/db/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

// ReadPlant reads a single plant in the database by finding it with the generated UUID
// from the database. Returns the &pb.Plant if found. else it returns an error
// if err code == 5 then the item was not found
func (s *PlantService) ReadPlant(ctx context.Context, in *pb.PlantId) (*pb.Plant, error) {
	log.Println("ReadPlant(server) invoked")
	queryPlant := Plant{
		Id: in.Id,
	}

	err := s.session.Collection("native.plants").Find(db.Cond{"id": queryPlant.Id}).One(&queryPlant)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, //return code [5]
			fmt.Sprintf("Error Not Found [%s]\n %s\n",
				queryPlant.Id,
				err.Error()))
	}
	// return protobuf
	return plantToProto(queryPlant), nil
}
