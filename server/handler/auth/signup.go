package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"

	"github.com/nepp-tumsat/documents-api/infrastructure"
	"github.com/nepp-tumsat/documents-api/infrastructure/persistence"
	"github.com/nepp-tumsat/documents-api/server/response"
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
			err = xerrors.Errorf("Error in uuid: %v", err)
			return
		}

		authRepo := persistence.NewAuthDB(infrastructure.DB)

		err = authRepo.InsertToUsers(userID.String(), requestBody.UserName)
		if err != nil {
			err = xerrors.Errorf("Error in repository: %v", err)
			return
		}

		response.Success(writer, "")
	}
}

type authSignUpRequest struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
