package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {

	user, err := UserDao.GetUser(0)

	assert.Nil(t, user, "We were not expecting a use with id 0")
	assert.NotNil(t, err, "We were expecting an error with user id is 0")
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode, "We were expecting 404 when user is not found")
	assert.EqualValues(t, "not found", err.Code, "We were expecing the correct erorr code")
	assert.EqualValues(t, "user 0 was not found", err.Message, "We were expecing the correct error message")

}

func TestGetUserIsFound(t *testing.T) {

	user, err := UserDao.GetUser(123)

	assert.NotNil(t, user, "We were expect a user to be found with id 123")
	assert.Nil(t, err, "We were expecing no error")

}
