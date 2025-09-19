package repository

import "microservice/internal/user/entity"

type UserRepository interface {
	GetAll() ([]*entity.User, error)
	Add(user *entity.User) (*entity.User, error)
}

type InMemoryUserRepo struct {
	users []*entity.User
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{users: []*entity.User{}}
}

func (r *InMemoryUserRepo) GetAll() ([]*entity.User, error) {
	return r.users, nil
}

func (r *InMemoryUserRepo) Add(user *entity.User) (*entity.User, error) {
	user.ID = int32(len(r.users) + 1)
	r.users = append(r.users, user)
	return user, nil
}
