package persistence

import (
	"database/sql"

	"github.com/nepp-tumsat/documents-api/infrastructure"
	"github.com/nepp-tumsat/documents-api/model"
	"golang.org/x/xerrors"
)

var db *sql.DB = infrastructure.DB

func SelectAll() ([]model.User, error) {
	rows, err := db.Query(`
    SELECT username
		FROM users;
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
