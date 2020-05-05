package persistence

import (
	"database/sql"

	"github.com/nepp-tumsat/documents-api/model"
	"golang.org/x/xerrors"
)

type AuthRepository interface {
	InsertToUsers(string, string) error
	InsertToUserAuths(string, string, string) error
	InsertToAuthTokens(string, string, string) error
	SelectUserAuthByEmail(string) (*model.UserAuth, error)
	SelectUserNameByUserID(string) (string, error)
	SelectUserIDByToken(string) (string, error)
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

func (a *authRepository) SelectUserAuthByEmail(email string) (*model.UserAuth, error) {
	row := a.db.QueryRow(`
		SELECT
		  user_id,
			email,
			hash
		FROM user_auths
		WHERE email=?;
	`, email)

	var userAuth model.UserAuth
	err := row.Scan(&userAuth.UserID, &userAuth.Email, &userAuth.Hash)
	if err != nil {
		err = xerrors.Errorf("Error in sql.DB: %v", err)
		return nil, err
	}
	return &userAuth, nil
}

func (a *authRepository) SelectUserNameByUserID(userID string) (string, error) {
	row := a.db.QueryRow(`
		SELECT
		  username
		FROM users
		WHERE id=?;
	`, userID)

	var username string
	err := row.Scan(&username)
	if err != nil {
		err = xerrors.Errorf("Error in sql.DB: %v", err)
		return "", err
	}
	return username, nil
}

func (a *authRepository) SelectUserIDByToken(token string) (string, error) {
	row := a.db.QueryRow(`
		SELECT
		  user_id
		FROM auth_tokens
		WHERE token=?;
	`, token)

	var userID string
	err := row.Scan(&userID)
	if err != nil {
		err = xerrors.Errorf("Error in sql.DB: %v", err)
		return "", err
	}
	return userID, nil
}
