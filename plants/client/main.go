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
	// list plants already within database
	listPlants(plantServiceClient)
	// insert an individual plant.
	plantId, _ := createPlant(plantServiceClient,
		&pb.Plant{
			CommonName: "False Rosemary",
			Genus:      "Conradina",
			Species:    "Canescens",
		})
	log.Printf("Plant with Id:%s, has been created\n",
		plantId.Id)
	time.Sleep(time.Second) // temporary wait time

	newPlantId := &pb.PlantId{
		Id: "8a2046b0-51ab-4fe2-9aa8-516cd5876ead", // the Id of the Seabeach-evening primrose
	}
	// should find the Seabeach-evening primrose
	secondPlantFromDatabase := readPlant(plantServiceClient, newPlantId)
	// found the seabeach evening primrose
	secondPlantFromDatabase.CommonName = "Primrose seabach evening"
	if _, err := updatePlant(plantServiceClient, secondPlantFromDatabase); err != nil {
		log.Println(err)
	}
	if err := deletePlant(plantServiceClient, plantId); err != nil {
		log.Println(err)
	} else {
		log.Printf("deleted\n %+v\n",
			plantId)
	}
}
