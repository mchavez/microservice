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
	return uc.repo.FindAll()
}

func (uc *UserUseCase) CreateUser(user *entity.User) (*entity.User, error) {
	return uc.repo.Save(user)
}

func (uc *UserUseCase) GetUserByID(id int64) (*entity.User, error) { // NEW
	return uc.repo.FindByID(id)
}

func (uc *UserUseCase) GetUsersByName(name string) ([]*entity.User, error) {
	return uc.repo.FindByName(name)
}
