package repository

import (
	"errors"
	"microservice/internal/user/entity"
)

type UserRepository interface {
	Save(user *entity.User) (*entity.User, error)
	FindAll() ([]*entity.User, error)
	FindByID(id int64) (*entity.User, error)        // NEW
	FindByName(name string) ([]*entity.User, error) // NEW
}

type InMemoryUserRepo struct {
	users []*entity.User
	ID    int64
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{users: []*entity.User{}}
}

func (r *InMemoryUserRepo) FindAll() ([]*entity.User, error) {
	return r.users, nil
}

func (r *InMemoryUserRepo) Save(user *entity.User) (*entity.User, error) {
	user.ID = int64(len(r.users) + 1)
	r.users = append(r.users, user)
	return user, nil
}

func (r *InMemoryUserRepo) FindByID(id int64) (*entity.User, error) {
	for _, u := range r.users {
		if u.ID == id {
			return u, nil
		}
	}

	return nil, errors.New("user not found")
}

func (r *InMemoryUserRepo) FindByName(name string) ([]*entity.User, error) {
	var results []*entity.User
	for _, u := range r.users {
		if u.Name == name {
			results = append(results, u)
		}
	}

	if len(results) == 0 {
		return nil, errors.New("no users found")
	}

	return results, nil
}
