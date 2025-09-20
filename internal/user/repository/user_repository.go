package repository

import "microservice/internal/user/entity"

type UserRepository interface {
	Save(user *entity.User) (*entity.User, error)
	FindAll() ([]*entity.User, error)
	FindByID(id int64) (*entity.User, error)
	FindByName(name string) ([]*entity.User, error) // NEW
}
