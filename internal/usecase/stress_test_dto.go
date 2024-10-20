package usecase

type ConfigDto struct {
	Url         string
	Requests    int
	Concurrency int
}

type ResultDto struct {
	TotalExecutionTimeInMilliseconds int64
	TotalRequests                    int
	SuccessFullRequests              map[int]int
	ErrorRequests                    map[int]int
}
