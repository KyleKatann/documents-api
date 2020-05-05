package persistence

import (
	"database/sql"

	"github.com/nepp-tumsat/documents-api/model"
	"golang.org/x/xerrors"
)

type AuthRepository interface {
	SelectAll() ([]model.User, error)
}

type authRepository struct {
	db *sql.DB
}

func NewAuthDB(db *sql.DB) AuthRepository {
	return &authRepository{db: db}
}

func (a *authRepository) SelectAll() ([]model.User, error) {
	rows, err := a.db.Query(`
    SELECT username FROM users;
	`)

	if err != nil {
		err = xerrors.Errorf("Error in sql.DB: %v", err)
		return nil, err
	}

	var users []model.User
	var user model.User
	for rows.Next() {
		err := rows.Scan(&user.UserName)
		if err != nil {
			err = xerrors.Errorf("Error in sql.DB: %v", err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
