package main

import (
	"fmt"
	"sync"

	"github.com/sjmillington/golang-microservices/github-api/src/api/utils/errors"

	"github.com/sjmillington/golang-microservices/github-api/src/api/domain/repositories"
	"github.com/sjmillington/golang-microservices/github-api/src/api/services"
)

var (
	success map[string]string
	fails   map[string]string
)

type createRepoResult struct {
	Request repositories.CreateRepoRequest
	Result  *repositories.CreateRepoResponse
	Error   errors.ApiError
}

func main() {
	requests := getRequests()

	fmt.Println(fmt.Sprintf("about to process %d requests", len(requests)))

	input := make(chan createRepoResult)
	buffer := make(chan bool, 10)

	var wg sync.WaitGroup
	//always pass a pointer so it's not a copy of the WG
	go handleResults(&wg, input)

	for _, request := range requests {
		buffer <- true
		wg.Add(1)
		go createRepo(request, input, buffer)
	}

	wg.Wait()
	close(input)

}

func getRequests() []repositories.CreateRepoRequest {

	results := make([]repositories.CreateRepoRequest, 0)

	for i := 0; i < 10000; i++ {
		results = append(results, repositories.CreateRepoRequest{Name: fmt.Sprintf("My new repo")})
	}
	return results
}

func createRepo(request repositories.CreateRepoRequest, input chan createRepoResult, buffer chan bool) {

	result, err := services.RepositoryService.CreateRepo(request)

	input <- createRepoResult{
		Request: request,
		Result:  result,
		Error:   err,
	}

	<-buffer

}

func handleResults(wg *sync.WaitGroup, input chan createRepoResult) {

	for result := range input {
		wg.Done()
		if result.Error != nil {
			fails[result.Request.Name] = result.Error.Message()
			continue
		}

		success[result.Request.Name] = result.Result.Name
	}

}
