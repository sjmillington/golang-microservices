package services

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/sjmillington/golang-microservices/github-api/src/api/utils/errors"

	"github.com/stretchr/testify/assert"

	"github.com/sjmillington/golang-microservices/github-api/src/api/domain/repositories"

	"github.com/sjmillington/golang-microservices/github-api/src/api/client/restclient"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidInputName(t *testing.T) {
	request := repositories.CreateRepoRequest{}

	result, err := RepositoryService.CreateRepo(request)

	assert.Nil(t, result)
	assert.NotNil(t, err)

	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "Invalid repository name", err.Message())

}

func TestCreateRepoErrorFromGitHub(t *testing.T) {
	restclient.FlushMockups()

	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message":"Requires Authentication"}`)),
		},
	})

	request := repositories.CreateRepoRequest{
		Name: "my-repo-name",
	}

	result, err := RepositoryService.CreateRepo(request)

	assert.Nil(t, result)
	assert.NotNil(t, err)

	assert.EqualValues(t, http.StatusUnauthorized, err.Status())
	assert.EqualValues(t, "Requires Authentication", err.Message())
}

func TestCreateRepoNoError(t *testing.T) {

	restclient.FlushMockups()

	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id":123}`)),
		},
	})

	request := repositories.CreateRepoRequest{
		Name: "my-repo-name",
	}

	result, err := RepositoryService.CreateRepo(request)

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.EqualValues(t, 123, result.Id)
	assert.EqualValues(t, "", result.Name)
	assert.EqualValues(t, "", result.Owner)

}

func TestCreateRepoConcurrentInvalidRequest(t *testing.T) {

	request := repositories.CreateRepoRequest{}

	output := make(chan repositories.CreateRepositoriesResult)

	service := repoService{}
	go service.createRepoConcurrent(request, output)

	result := <-output

	assert.NotNil(t, result)
	assert.Nil(t, result.Response)
	assert.NotNil(t, result.Error)
	assert.EqualValues(t, "Invalid repository name", result.Error.Message())
	assert.EqualValues(t, http.StatusBadRequest, result.Error.Status())
}

func TestCreateRepoConcurrentErrorFromGithub(t *testing.T) {

	restclient.FlushMockups()

	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message":"Requires Authentication"}`)),
		},
	})

	request := repositories.CreateRepoRequest{
		Name: "my-repo-name",
	}

	output := make(chan repositories.CreateRepositoriesResult)

	service := repoService{}
	go service.createRepoConcurrent(request, output)

	result := <-output

	assert.NotNil(t, result)
	assert.Nil(t, result.Response)
	assert.NotNil(t, result.Error)
	assert.EqualValues(t, "Requires Authentication", result.Error.Message())
	assert.EqualValues(t, http.StatusUnauthorized, result.Error.Status())
}

func TestCreateRepoConcurrentNoError(t *testing.T) {

	restclient.FlushMockups()

	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id":123}`)),
		},
	})

	request := repositories.CreateRepoRequest{
		Name: "my-repo-name",
	}

	output := make(chan repositories.CreateRepositoriesResult)

	service := repoService{}
	go service.createRepoConcurrent(request, output)

	result := <-output

	assert.NotNil(t, result)
	assert.NotNil(t, result.Response)
	assert.Nil(t, result.Error)
	assert.EqualValues(t, 123, result.Response.Id)
	assert.EqualValues(t, "", result.Response.Name)
}

func TestHandleRepoResults(t *testing.T) {
	input := make(chan repositories.CreateRepositoriesResult)
	output := make(chan repositories.CreateReposResponse)
	var wg sync.WaitGroup

	service := repoService{}
	go service.handleRepoResults(&wg, input, output)

	wg.Add(1)
	go func() {
		input <- repositories.CreateRepositoriesResult{
			Error: errors.NewBadRequestAPIError("Invalid repository name"),
		}
	}()

	wg.Wait()
	close(input)
	result := <-output

	assert.NotNil(t, result)
	assert.EqualValues(t, 0, result.StatusCode)
	assert.EqualValues(t, 1, len(result.Results))

	assert.NotNil(t, result.Results[0].Error)
	assert.EqualValues(t, http.StatusBadRequest, result.Results[0].Error.Status())
	assert.EqualValues(t, "Invalid repository name", result.Results[0].Error.Message())

}

func TestCreateReposInvalidRequests(t *testing.T) {

	requests := []repositories.CreateRepoRequest{
		{},
		{Name: "      "},
	}

	result, err := RepositoryService.CreateRepos(requests)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, 2, len(result.Results))
	assert.EqualValues(t, http.StatusBadRequest, result.StatusCode)

	assert.Nil(t, result.Results[0].Response)
	assert.EqualValues(t, http.StatusBadRequest, result.Results[0].Error.Status())
	assert.EqualValues(t, "Invalid repository name", result.Results[0].Error.Message())
	assert.Nil(t, result.Results[1].Response)
	assert.EqualValues(t, http.StatusBadRequest, result.Results[1].Error.Status())
	assert.EqualValues(t, "Invalid repository name", result.Results[1].Error.Message())

}

func TestCreateReposPartialValidRequests(t *testing.T) {

	restclient.FlushMockups()

	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id":123, "name": "testing", "owner": {"login": "sjmillington"}}`)),
		},
	})

	requests := []repositories.CreateRepoRequest{
		{},
		{Name: "my valid repo name"},
	}

	result, err := RepositoryService.CreateRepos(requests)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, 2, len(result.Results))
	assert.EqualValues(t, http.StatusPartialContent, result.StatusCode)

	assert.Nil(t, result.Results[0].Response)
	assert.EqualValues(t, "Invalid repository name", result.Results[0].Error.Message())
	assert.EqualValues(t, http.StatusBadRequest, result.Results[0].Error.Status())

	assert.Nil(t, result.Results[1].Error)
	assert.EqualValues(t, 123, result.Results[1].Response.Id)
	assert.EqualValues(t, "testing", result.Results[1].Response.Name)
	assert.EqualValues(t, "sjmillington", result.Results[1].Response.Owner)

}
