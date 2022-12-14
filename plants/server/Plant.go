package main

import pb "github.com/mbcarruthers/gRPCplants/plants/proto"

// Plant exists to be an intermediary type between the database & the go code as well
// as to avoid copying the protobuf implementation.
// A good explanation of the latter can be found in the answers to this stackoverflow question.
// https://stackoverflow.com/questions/64183794/why-do-the-go-generated-protobuf-files-contain-mutex-locks
type Plant struct {
	Id         string `db:"id,omitempty"` // Todo: Id should be lowercase, but changing it throughout the files breaks the update function somehow.
	CommonName string `db:"common_name"`
	Genus      string `db:"genus"`
	Species    string `db:"species"`
}

// plantToProto: Converts a Plant type to gRPC.Plant to
// avoid shallow copying the protocol implementation
// when sending/streaming a response to the client.
func plantToProto(p Plant) *pb.Plant {
	return &pb.Plant{
		Id:         p.Id,
		CommonName: p.CommonName,
		Genus:      p.Genus,
		Species:    p.Species,
	}
}

// protoToPlant is a utility function to change a protobuf plant
// to a database compatable plant.
func protoToPlant(p *pb.Plant) *Plant {
	return &Plant{
		Id:         p.Id,
		CommonName: p.CommonName,
		Genus:      p.Genus,
		Species:    p.Species,
	}
}
