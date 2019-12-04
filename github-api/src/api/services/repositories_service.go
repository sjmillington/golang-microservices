package services

import (
	"strings"

	"github.com/sjmillington/golang-microservices/github-api/src/api/config"
	"github.com/sjmillington/golang-microservices/github-api/src/api/providers/github_provider"

	"github.com/sjmillington/golang-microservices/github-api/src/api/domain/github"

	"github.com/sjmillington/golang-microservices/github-api/src/api/domain/repositories"
	"github.com/sjmillington/golang-microservices/github-api/src/api/utils/errors"
)

type repoService struct{}

type repoServiceInterface interface {
	CreateRepo(req repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
}

var (
	RepositoryService repoServiceInterface
)

func init() {
	RepositoryService = &repoService{}
}

func (r *repoService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	input.Name = strings.TrimSpace(input.Name)

	if input.Name == "" {
		return nil, errors.NewBadRequestAPIError("Invalid repository name")
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
