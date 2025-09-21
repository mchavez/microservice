package usecase_test

import (
	"microservice/internal/user/entity"
	"microservice/internal/user/repository"
	"microservice/internal/user/usecase"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestUserUseCase_AddAndGetUsers(t *testing.T) {
	logger := logrus.New()
	repo := repository.NewInMemoryUserRepo()
	uc := usecase.NewUserUseCase(repo, logger)

	// Add a user
	user := &entity.User{Name: "Alice"}
	created, err := uc.CreateUser(user)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if created.ID == 0 {
		t.Errorf("expected non-zero ID, got %d", created.ID)
	}

	// Fetch all users
	users, err := uc.GetUsers()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(users) != 1 {
		t.Errorf("expected 1 user, got %d", len(users))
	}
	if users[0].Name != "Alice" {
		t.Errorf("expected name Alice, got %s", users[0].Name)
	}
}
