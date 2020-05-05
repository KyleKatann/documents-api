package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"

	"github.com/nepp-tumsat/documents-api/infrastructure"
	"github.com/nepp-tumsat/documents-api/infrastructure/persistence"
	"github.com/nepp-tumsat/documents-api/server/response"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/xerrors"
)

func HandleAuthSignUp() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		var requestBody authSignUpRequest
		err := json.NewDecoder(request.Body).Decode(&requestBody)
		if err != nil {
			log.Printf("%+v\n", xerrors.Errorf("Error in json: %v", err))
			response.BadRequest(writer, "Can't decode of json")
			return
		}

		userID, err := uuid.NewRandom()
		if err != nil {
			log.Printf("%+v\n", xerrors.Errorf("Error in uuid: %v", err))
			return
		}

		authRepo := persistence.NewAuthDB(infrastructure.DB)

		err = authRepo.InsertToUsers(userID.String(), requestBody.UserName)
		if err != nil {
			log.Printf("%+v\n", xerrors.Errorf("Error in repository: %v", err))
			return
		}

		hash, err := passwordToHash(requestBody.Password)
		if err != nil {
			log.Printf("%+v\n", xerrors.Errorf("Error in bcrypt: %v", err))
			return
		}

		err = authRepo.InsertToUserAuths(userID.String(), requestBody.Email, hash)
		if err != nil {
			log.Printf("%+v\n", xerrors.Errorf("Error in repository: %v", err))
			return
		}

		authTokenID, err := uuid.NewRandom()
		if err != nil {
			log.Printf("%+v\n", xerrors.Errorf("Error in uuid: %v", err))
			return
		}

		token, err := uuid.NewRandom()
		if err != nil {
			log.Printf("%+v\n", xerrors.Errorf("Error in uuid: %v", err))
			return
		}

		err = authRepo.InsertToAuthTokens(authTokenID.String(), userID.String(), token.String())
		if err != nil {
			log.Printf("%+v\n", xerrors.Errorf("Error in repository: %v", err))
			return
		}

		response.Success(writer, authSignUpResponse{UserName: requestBody.UserName, Token: token.String()})
	}
}

func passwordToHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hash), err
}

type authSignUpRequest struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type authSignUpResponse struct {
	UserName string `json:"username"`
	Token    string `json:"token"`
}
