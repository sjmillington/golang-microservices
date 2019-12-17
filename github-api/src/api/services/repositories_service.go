package services

import (
	"net/http"
	"sync"

	"github.com/sjmillington/golang-microservices/github-api/src/api/config"
	"github.com/sjmillington/golang-microservices/github-api/src/api/providers/github_provider"

	"github.com/sjmillington/golang-microservices/github-api/src/api/domain/github"

	"github.com/sjmillington/golang-microservices/github-api/src/api/domain/repositories"
	"github.com/sjmillington/golang-microservices/github-api/src/api/utils/errors"
)

type repoService struct{}

type repoServiceInterface interface {
	CreateRepo(req repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	CreateRepos(req []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError)
}

var (
	RepositoryService repoServiceInterface
)

func init() {
	RepositoryService = &repoService{}
}

func (r *repoService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {

	if err := input.Validate(); err != nil {
		return nil, err
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)

	if err != nil {
		return nil, errors.NewAPIError(err.StatusCode, err.Message)
	}

	result := repositories.CreateRepoResponse{
		Id:    response.ID,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}

	return &result, nil
}

func (r *repoService) CreateRepos(request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError) {

	input := make(chan repositories.CreateRepositoriesResult)
	output := make(chan repositories.CreateReposResponse)
	defer close(output)

	var wg sync.WaitGroup
	go r.handleRepoResults(&wg, input, output)

	for _, current := range request {
		wg.Add(1)
		go r.createRepoConcurrent(current, input)
	}

	wg.Wait()
	close(input)

	result := <-output

	creations := 0

	for _, current := range result.Results {
		if current.Response != nil {
			creations++
		}
	}

	if creations == 0 {
		result.StatusCode = result.Results[0].Error.Status()
	} else if creations == len(request) {
		result.StatusCode = http.StatusCreated
	} else {
		result.StatusCode = http.StatusPartialContent
	}

	return result, nil

}

func (r *repoService) handleRepoResults(wg *sync.WaitGroup, input chan repositories.CreateRepositoriesResult, output chan repositories.CreateReposResponse) {

	var results repositories.CreateReposResponse

	for incEvent := range input {

		repoResult := repositories.CreateRepositoriesResult{
			Response: incEvent.Response,
			Error:    incEvent.Error,
		}

		results.Results = append(results.Results, repoResult)
		wg.Done()
	}

	output <- results

}

func (r *repoService) createRepoConcurrent(input repositories.CreateRepoRequest, output chan repositories.CreateRepositoriesResult) {
	if err := input.Validate(); err != nil {
		output <- repositories.CreateRepositoriesResult{Error: err}
		return
	}

	result, err := r.CreateRepo(input)

	if err != nil {
		output <- repositories.CreateRepositoriesResult{Error: err}
		return
	}

	output <- repositories.CreateRepositoriesResult{Response: result}

}
