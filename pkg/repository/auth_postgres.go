package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sdf0106/os-project"
	"time"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user project.User) (project.User, error) {
	query := fmt.Sprintf("INSERT  INTO %s (name, email, password_hash, created_at) values ($1, $2, $3, $4) RETURNING ID", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Email, user.Password, time.Now())
	if err := row.Scan(&user.Id); err != nil {
		return user, err
	}
	return user, nil
}

func (r *AuthPostgres) GetUser(email, password string) (project.User, error) {
	var user project.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, email, password)

	return user, err
}
