package usecase

import (
	"fmt"
	"go-stress-tester/internal/infra"
	"strings"
	"sync"
	"time"
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

	startTime := time.Now() // Start time

	var mu sync.Mutex
	var wg sync.WaitGroup
	sem := make(chan struct{}, config.Concurrency)
	progress := 0
	barLength := 50

	for i := 0; i < config.Requests; i++ {
		wg.Add(1)
		sem <- struct{}{} // Acquire a slot
		go func() {
			defer wg.Done()
			code, err := r.networkClient.DoRequest(config.Url)
			mu.Lock()
			if err != nil || code < 200 || code >= 300 {
				result.ErrorRequests[code]++
			} else {
				result.SuccessFullRequests[code]++
			}
			progress++
			if progress%10 == 0 || progress == config.Requests {
				percentage := float64(progress) / float64(config.Requests)
				bar := int(percentage * float64(barLength))
				fmt.Printf("\rProgress: [%s%s] %d/%d", strings.Repeat("=", bar), strings.Repeat(" ", barLength-bar), progress, config.Requests)
			}
			mu.Unlock()
			<-sem // Release the slot
		}()
	}

	wg.Wait()
	result.TotalExecutionTimeInMilliseconds = time.Since(startTime).Milliseconds()
	result.TotalRequests = progress
	return result
}
