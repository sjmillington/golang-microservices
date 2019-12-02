package services

import (
	"net/http"
	"testing"

	"github.com/sjmillington/golang-microservices/mvc/domain"
	"github.com/sjmillington/golang-microservices/mvc/utils"

	"github.com/stretchr/testify/assert"
)

var (
	UserDaoMock userDaoMock

	getUserFunction func(userId int64) (*domain.User, *utils.ApplicationError)
)

func init() {
	domain.UserDao = &userDaoMock{}
}

type userDaoMock struct{}

func (m *userDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userId)
}

func TestGetUserNotFoundInDatabase(t *testing.T) {

	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			Message:    "user 0 was not found",
			StatusCode: http.StatusNotFound,
		}
	}

	user, err := UsersService.GetUser(0)

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 was not found", err.Message)
}

func TestGetUserFoundInDatabase(t *testing.T) {

	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			ID: uint64(userId),
		}, nil
	}

	user, err := UsersService.GetUser(123)

	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.EqualValues(t, 123, user.ID)

}
