package repository_test

import (
	"microservice/internal/user/entity"
	"microservice/internal/user/repository"
	"testing"
)

func TestInMemoryUserRepo(t *testing.T) {
	repo := repository.NewInMemoryUserRepo()

	// Add user
	u, err := repo.Add(&entity.User{Name: "Bob"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if u.ID != 1 {
		t.Errorf("expected ID 1, got %d", u.ID)
	}

	// Get all
	users, err := repo.GetAll()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(users) != 1 {
		t.Errorf("expected 1 user, got %d", len(users))
	}
}
