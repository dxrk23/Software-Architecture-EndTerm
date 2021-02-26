package computeService

import (
	"context"
	"dimash/endterm/pkg/api/avgAPI"
	"dimash/endterm/pkg/api/primeAPI"
)

// gRPC server ...
type GRPCServer struct {}

func (s *GRPCServer) Decomposite(ctx context.Context, req *primeAPI.DecompositeRequest) (*primeAPI.DecompositeResponse, error){

	n := req.GetNumber()
	var result []int32

	for n%2 == 0 {
		result = append(result, 2)
		n = n / 2
	}

	var i int32

	for i = 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			result = append(result, i)
			n = n / i
		}
	}

	if n > 2 {
		result = append(result, n)
	}

	return &primeAPI.DecompositeResponse{Results: result}, nil
}

func (s *GRPCServer) ComputeAverage(ctx context.Context, req *avgAPI.AverageRequest) (*avgAPI.AverageResponse, error){
	var result float32 = 0
	var nums []int32 = req.GetNumbers()

	for _, number:= range nums {
		result = result + float32(number)
	}
	result = result / float32(len(nums))

	return &avgAPI.AverageResponse{Result: result}, nil
}
