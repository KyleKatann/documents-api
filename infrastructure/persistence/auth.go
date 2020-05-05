package persistence

import (
	"database/sql"

	"golang.org/x/xerrors"
)

type AuthRepository interface {
	InsertToUsers(string, string) error
	InsertToUserAuths(string, string, string) error
	InsertToAuthTokens(string, string, string) error
	SelectHashByEmail(string) (string, error)
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

func (a *authRepository) InsertToUserAuths(userID, email, hash string) error {
	stmt, err := a.db.Prepare(`
		INSERT INTO
			user_auths(
			  user_id,
			  email,
				hash
			)
		VALUES(?,?,?);
	`)
	if err != nil {
		err = xerrors.Errorf("Error in sql.DB: %v", err)
		return err
	}

	_, err = stmt.Exec(userID, email, hash)
	if err != nil {
		err = xerrors.Errorf("Error in sql.DB: %v", err)
		return err
	}
	return nil
}

func (a *authRepository) InsertToAuthTokens(authTokenID, userID, token string) error {
	stmt, err := a.db.Prepare(`
		INSERT INTO
			auth_tokens(
			  id,
			  user_id,
				token
			)
		VALUES(?,?,?);
	`)
	if err != nil {
		err = xerrors.Errorf("Error in sql.DB: %v", err)
		return err
	}

	_, err = stmt.Exec(authTokenID, userID, token)
	if err != nil {
		err = xerrors.Errorf("Error in sql.DB: %v", err)
		return err
	}
	return nil
}

func (a *authRepository) SelectHashByEmail(email string) (string, error) {
	row := a.db.QueryRow(`
		SELECT
			hash
		FROM user_auths
		WHERE email=?;
	`, email)

	var hash string
	err := row.Scan(&hash)
	if err != nil {
		err = xerrors.Errorf("Error in sql.DB: %v", err)
		return "", err
	}
	return hash, nil
}
