package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"dimash/endterm/pkg/api/primeAPI"
	"dimash/endterm/pkg/computeService"
)

func main() {
	s := grpc.NewServer()
	srv := &computeService.GRPCServer{}

	primeAPI.RegisterPrimeDecompositorServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil{
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil{
		log.Fatal(err)
	}
}
