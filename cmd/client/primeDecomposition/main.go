package main

import (
	"context"
	"dimash/endterm/pkg/api/primeAPI"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main(){

	var temp int
	fmt.Print("Enter number: ")
	_, err := fmt.Scanf("%d", &temp)

	number := int32(temp)

	conn , err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := primeAPI.NewPrimeDecompositorClient(conn)
	res, err := c.Decomposite(context.Background(), &primeAPI.DecompositeRequest{Number: int32(number)})

	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.GetResults())
}