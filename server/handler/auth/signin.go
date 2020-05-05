package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nepp-tumsat/documents-api/infrastructure"
	"github.com/nepp-tumsat/documents-api/infrastructure/persistence"
	"github.com/nepp-tumsat/documents-api/server/response"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/xerrors"
)

func HandleAuthSignIn() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		var requestBody authSignInRequest
		err := json.NewDecoder(request.Body).Decode(&requestBody)
		if err != nil {
			log.Printf("%+v\n", xerrors.Errorf("Error in json: %v", err))
			response.BadRequest(writer, "Can't decode of json")
			return
		}

		authRepo := persistence.NewAuthDB(infrastructure.DB)

		hash, err := authRepo.SelectHashByEmail(requestBody.Email)
		if err != nil {
			log.Printf("%+v\n", xerrors.Errorf("Error in repository: %v", err))
			return
		}

		err = passwordVerify(requestBody.Password, hash)
		if err != nil {
			log.Printf("%+v\n", xerrors.Errorf("Error in request: %v", err))
			response.BadRequest(writer, "Can't verify of password")
			return
		}

		response.Success(writer, hash)
	}

}

func passwordVerify(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

type authSignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
