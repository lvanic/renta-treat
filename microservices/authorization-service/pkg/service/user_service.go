package service

import (
	"crypto/rand"
	"errors"
	"go/pkg/model"
	"go/pkg/repository"
	"go/pkg/utils"
	"sync"
)

const SALT_LENGTH = 16

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
	randomBytes := make([]byte, SALT_LENGTH)
	if _, err := rand.Read(randomBytes); err != nil {
		panic(err)
	}
	user.Salt = randomBytes
	user.Password = utils.HashPassword(user.Password, user.Salt)
	err := s.userRepository.Create(user)
	return err
}

func (s *UserService) AuthorizeUser(user *model.User) (*model.User, error) {
	userHandler, err := s.userRepository.GetOneByEmail(user.Email)
	if err != nil {
		return nil, err
	}
	if isPasswordsMatch := utils.ComparePassword(user.Password, userHandler.Salt, userHandler.Password); isPasswordsMatch {
		return userHandler, nil
	} else {
		return nil, errors.New("Incorrect credetinals")
	}
}

func (s *UserService) GetUserById(id int) (*model.User, error) {
	user, err := s.userRepository.GetOneById(id)
	return user, err
}
