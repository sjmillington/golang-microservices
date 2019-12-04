package services

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

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
