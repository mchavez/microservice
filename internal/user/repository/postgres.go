package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"microservice/internal/user/entity"

	_ "github.com/lib/pq"
)

type PostgresUserRepo struct {
	db *sql.DB
}

// NewPostgresUserRepo initializes a new PostgresUserRepo instance.
func NewPostgresUserRepo(host, port, user, password, dbname string) (*PostgresUserRepo, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresUserRepo{db: db}, nil
}

// FindAll retrieves all users from the database.
func (r *PostgresUserRepo) FindAll() ([]*entity.User, error) {
	rows, err := r.db.Query("SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		var u entity.User
		if err := rows.Scan(&u.ID, &u.Name); err != nil {
			return nil, err
		}
		users = append(users, &u)
	}
	return users, nil
}

// Save inserts a new user into the database and returns the user with the generated ID.
func (r *PostgresUserRepo) Save(user *entity.User) (*entity.User, error) {
	err := r.db.QueryRow(
		"INSERT INTO users (name) VALUES ($1) RETURNING id",
		user.Name,
	).Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// FindByID retrieves a user by their ID.
func (r *PostgresUserRepo) FindByID(id int64) (*entity.User, error) {
	var u entity.User
	err := r.db.QueryRow("SELECT id, name FROM users WHERE id = $1", id).Scan(&u.ID, &u.Name)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// FindByName retrieves users by their name.
func (r *PostgresUserRepo) FindByName(name string) ([]*entity.User, error) {
	rows, err := r.db.Query("SELECT id, name FROM users WHERE name = $1", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		var u entity.User
		if err := rows.Scan(&u.ID, &u.Name); err != nil {
			return nil, err
		}
		users = append(users, &u)
	}

	if len(users) == 0 {
		return nil, errors.New("no users found")
	}
	return users, nil
}
