package auth

import (
	"net/http"

	"github.com/nepp-tumsat/documents-api/server/response"
)

func HandleAuthSignUp() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		response.Success(writer, "success")
	}
}

type authSignUpRequest struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
