//go:build integration
// +build integration

package repository_test

import (
	"microservice/internal/user/entity"
	"microservice/internal/user/repository"
	"os"
	"testing"
)

func TestPostgresUserRepo(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	pass := getEnv("DB_PASS", "postgres")
	dbname := getEnv("DB_NAME", "usersdb")

	repo, err := repository.NewPostgresUserRepo(host, port, user, pass, dbname)
	if err != nil {
		t.Fatalf("could not connect to postgres: %v", err)
	}

	// Add user
	u, err := repo.Save(&entity.User{Name: "Charlie"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if u.ID == 0 {
		t.Errorf("expected non-zero ID")
	}

	// Get all
	users, err := repo.FindAll()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(users) == 0 {
		t.Errorf("expected at least 1 user")
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
