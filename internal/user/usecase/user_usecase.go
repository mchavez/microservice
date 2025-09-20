package usecase

import (
	"microservice/internal/user/entity"
	"microservice/internal/user/repository"

	"github.com/sirupsen/logrus"
)

type UserUseCase struct {
	repo   repository.UserRepository
	logger *logrus.Logger
}

func NewUserUseCase(r repository.UserRepository, l *logrus.Logger) *UserUseCase {
	return &UserUseCase{repo: r, logger: l}
}

func (uc *UserUseCase) GetUsers() ([]*entity.User, error) {
	return uc.repo.FindAll()
}

func (uc *UserUseCase) CreateUser(user *entity.User) (*entity.User, error) {
	uc.logger.WithField("name", user.Name).Info("creating user")

	user, err := uc.repo.Save(user)
	if err != nil {
		uc.logger.WithError(err).Error("failed to create user")
		return nil, err
	}

	uc.logger.WithField("id", user.ID).Info("user created successfully")
	return user, nil
}

func (uc *UserUseCase) GetUserByID(id int64) (*entity.User, error) {
	return uc.repo.FindByID(id)
}

func (uc *UserUseCase) GetUsersByName(name string) ([]*entity.User, error) {
	return uc.repo.FindByName(name)
}
