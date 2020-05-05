package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nepp-tumsat/documents-api/server/response"
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

		response.Success(writer, requestBody)
	}
}

type authSignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
