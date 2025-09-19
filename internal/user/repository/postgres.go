package repository

import (
	"database/sql"
	"fmt"
	"microservice/internal/user/entity"

	_ "github.com/lib/pq"
)

type PostgresUserRepo struct {
	db *sql.DB
}

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

func (r *PostgresUserRepo) GetAll() ([]*entity.User, error) {
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

func (r *PostgresUserRepo) Add(user *entity.User) (*entity.User, error) {
	err := r.db.QueryRow(
		"INSERT INTO users (name) VALUES ($1) RETURNING id",
		user.Name,
	).Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
