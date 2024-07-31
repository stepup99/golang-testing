package service

import "github.com/stepup99/golang-testing/user_gin/domain"

type UserService interface {
	GetUser(userid int) (*domain.User, error)
}
