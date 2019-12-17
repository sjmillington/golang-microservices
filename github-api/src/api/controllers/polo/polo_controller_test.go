package polo

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sjmillington/golang-microservices/github-api/src/api/utils/test_utils"
	"github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "polo", polo)
}

func TestPolo(t *testing.T) {

	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/marco", nil)

	c := test_utils.GetMockedContext(request, response)

	Polo(c)

	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.EqualValues(t, "polo", response.Body.String())

}
