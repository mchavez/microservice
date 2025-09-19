package usecase

import (
	"microservice/internal/user/entity"
	"microservice/internal/user/repository"
)

type UserUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(r repository.UserRepository) *UserUseCase {
	return &UserUseCase{repo: r}
}

func (uc *UserUseCase) GetUsers() ([]*entity.User, error) {
	return uc.repo.GetAll()
}

func (uc *UserUseCase) CreateUser(user *entity.User) (*entity.User, error) {
	return uc.repo.Add(user)
}
