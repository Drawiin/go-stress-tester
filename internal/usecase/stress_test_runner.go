package usecase

import (
	"fmt"
	"go-stress-tester/internal/infra"
)

type StressTestRunner struct {
	networkClient infra.NetworkClient
}

func NewStressTestRunner(networkClient infra.NetworkClient) StressTestRunner {
	return StressTestRunner{networkClient: networkClient}
}

func (r *StressTestRunner) RunStressTest(config ConfigDto) ResultDto {
	fmt.Println("run stress test with config", config)
	result := ResultDto{
		ErrorRequests:       make(map[int]int),
		SuccessFullRequests: make(map[int]int),
	}
	for i := 0; i < config.Requests; i++ {
		code, err := r.networkClient.DoRequest(config.Url)
		if err != nil {
			result.ErrorRequests[code]++
		} else {
			result.SuccessFullRequests[code]++
		}
	}
	return result
}
