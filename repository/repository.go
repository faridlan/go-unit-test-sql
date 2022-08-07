package repository

import "github.com/faridlan/go-unit-test-sql/model"

type Repository interface {
	Close()
	FindById(id string) (*model.User, error)
	FindAll() ([]*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(id string) error
}
