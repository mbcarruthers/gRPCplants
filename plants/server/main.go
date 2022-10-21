package main

import (
	pb "github.com/mbcarruthers/gRPCplants/plants/proto"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/cockroachdb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type PlantService struct {
	pb.PlantServiceServer
	session db.Session
}

// Todo: Create this into a singleton pattern , okay for now though.
func NewPlantService(ses db.Session) *PlantService {
	return &PlantService{
		session: ses,
	}
}

var (
	settings cockroachdb.ConnectionURL
)

const (
	addr = "0.0.0.0:50051"
)

func init() {
	// initial database connection set-up
	settings = cockroachdb.ConnectionURL{
		Host:     "localhost:26257",
		Database: "native",
		User:     "root",
		Options:  map[string]string{"sslmode": "disable"},
	}
}

func main() {
	// set up TCP connection for gRPC server
	if list, err := net.Listen("tcp", addr); err != nil {
		log.Fatalf("Error listening to %s \n %s\n",
			addr,
			err.Error())
	} else {
		log.Println("Listening at ", addr)
		// set up session for cockroach database
		session, err := cockroachdb.Open(settings)
		if err != nil {
			log.Fatal("cockroachdb.Open: ", err)
		}
		// create new gRPC service for database
		plantService := NewPlantService(session)
		defer plantService.session.Close()
		defer func() {
			if r := recover(); r != nil {
				log.Println("[Panic]:  %+v \n", r)
			}
		}()

		log.Println("Connected to Cockroachdb")

		s := grpc.NewServer()
		// register database service to gRPC server
		pb.RegisterPlantServiceServer(s, plantService)

		// listen for connections from client(s)
		if err = s.Serve(list); err != nil {
			log.Fatalf("gRPC listen error \n %s\n",
				err.Error())
		}
	}
}
