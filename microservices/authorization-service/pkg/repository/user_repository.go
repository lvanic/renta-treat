package repository

import (
	"errors"
	"go/pkg/model"
	"sync"

	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

var once sync.Once
var instance *UserRepository

func NewUserRepository(db *gorm.DB) *UserRepository {
	once.Do(func() {
		instance = &UserRepository{Db: db}
	})
	return instance
}

func (r *UserRepository) Create(user *model.User) error {
	if err := r.Db.Create(user); err != nil {
		return errors.New("Create user error")
	} else {
		return nil
	}
}

func (r *UserRepository) GetOne(id int) (*model.User, error) {

	user := model.User{}
	if result := r.Db.Find(&user, id); result.Error != nil {
		return nil, errors.New("UserNotFound")
	} else if user.ID == 0 {
		return nil, errors.New("UserNotFound")
	} else {
		return &user, nil
	}
}
