package main

import (
	"bufio"
	"context"
	"dimash/endterm/pkg/api/avgAPI"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	"google.golang.org/grpc"
)

func main(){
	var nums []int32
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter numbers: ")
	numbers, _ := reader.ReadString('\n')

	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	submatchall := re.FindAllString(numbers, -1)

	for _, element := range submatchall {
		temp,err := strconv.ParseInt(element,10,32)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, int32(temp))
	}

	conn , err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := avgAPI.NewComputeAverageClient(conn)
	res, err := c.ComputeAverage(context.Background(), &avgAPI.AverageRequest{Numbers: nums})

	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.GetResult())
}