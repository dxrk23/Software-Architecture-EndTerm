package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"dimash/endterm/pkg/api/avgAPI"
	"dimash/endterm/pkg/computeService"
)

func main() {
	s := grpc.NewServer()
	srv := &computeService.GRPCServer{}

	avgAPI.RegisterComputeAverageServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil{
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil{
		log.Fatal(err)
	}
}
