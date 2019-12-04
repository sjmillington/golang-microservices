package repositories

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/sjmillington/golang-microservices/github-api/src/api/utils/test_utils"

	"github.com/sjmillington/golang-microservices/github-api/src/api/domain/repositories"

	"github.com/sjmillington/golang-microservices/github-api/src/api/client/restclient"
	"github.com/sjmillington/golang-microservices/github-api/src/api/utils/errors"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidJSONRequest(t *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(``))

	c := test_utils.GetMockedContext(request, response)

	CreateRepo(c)

	apiErr, err := errors.NewApiErrFromBytes(response.Body.Bytes())

	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(t, "invalid json body", apiErr.Message())

}

func TestCreateRepoErrorFromGithub(t *testing.T) {

	restclient.FlushMockups()

	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message":"Requires Authentication"}`)),
		},
	})

	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name":"hello-repo"}`))

	c := test_utils.GetMockedContext(request, response)

	CreateRepo(c)

	apiErr, err := errors.NewApiErrFromBytes(response.Body.Bytes())

	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusUnauthorized, apiErr.Status())
	assert.EqualValues(t, "Requires Authentication", apiErr.Message())

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

	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name":"hello-repo"}`))

	c := test_utils.GetMockedContext(request, response)

	CreateRepo(c)

	var result repositories.CreateRepoResponse

	err := json.Unmarshal(response.Body.Bytes(), &result)

	assert.Nil(t, err)
	assert.EqualValues(t, 123, result.Id)

}
