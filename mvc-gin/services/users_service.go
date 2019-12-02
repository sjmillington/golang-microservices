package services

import (
	"github.com/sjmillington/golang-microservices/mvc-gin/domain"
	"github.com/sjmillington/golang-microservices/mvc-gin/utils"
)

type usersService struct {
}

var (
	UsersService usersService
)

func (u *usersService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(userId)
}
