package main

import (
	pb "github.com/mbcarruthers/gRPCplants/plants/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const (
	addr = "localhost:50051"
)

func main() {
	// try to established connection to gRPC server
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to %s\n %s\n",
			addr,
			err.Error())
	}
	defer conn.Close()

	log.Printf("Connection established to %s \n", addr)
	time.Sleep(time.Second) // temporary wait time for now
	// create client service for accessing gRPC service to cockroach database.
	plantServiceClient := pb.NewPlantServiceClient(conn)
	listPlants(plantServiceClient)

}
