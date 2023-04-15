package service

import (
	"go/pkg/model"
	"go/pkg/repository"
	"sync"
)

type UserService struct {
	userRepository repository.UserRepository
}

var once sync.Once
var instance *UserService

func NewUserService(userRepository repository.UserRepository) *UserService {
	once.Do(func() {
		instance = &UserService{userRepository}
	})
	return instance
}

func (s *UserService) CreateUser(user *model.User) error {
	if err := s.userRepository.Create(user); err != nil {
		return err
	} else {
		return nil
	}
}

func (s *UserService) GetUserById(id int) (*model.User, error) {

	if user, err := s.userRepository.GetOne(id); err != nil {
		return nil, err
	} else {
		return user, nil
	}
}
