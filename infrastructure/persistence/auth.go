package persistence

import (
	"database/sql"

	"golang.org/x/xerrors"
)

type AuthRepository interface {
	InsertToUsers(string, string) error
}

type authRepository struct {
	db *sql.DB
}

func NewAuthDB(db *sql.DB) AuthRepository {
	return &authRepository{db: db}
}

func (a *authRepository) InsertToUsers(userID, userName string) error {
	stmt, err := a.db.Prepare(`
		INSERT INTO
			users(
			  id,
			  username
			)
		VALUES(?,?);
	`)
	if err != nil {
		err = xerrors.Errorf("Error in sql.DB: %v", err)
		return err
	}

	_, err = stmt.Exec(userID, userName)
	if err != nil {
		err = xerrors.Errorf("Error in sql.DB: %v", err)
		return err
	}
	return nil
}
