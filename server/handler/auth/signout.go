package auth

import (
	"log"
	"net/http"

	"github.com/nepp-tumsat/documents-api/dcontext"
	"github.com/nepp-tumsat/documents-api/server/response"
	"golang.org/x/xerrors"
)

func HandleAuthSignOut() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		ctx := request.Context()
		userID := dcontext.GetUserIDFromContext(ctx)
		if len(userID) == 0 {
			log.Printf("%+v\n", xerrors.New("Error in dcontext"))
			response.InternalServerError(writer, "Can't get userID")
			return
		}

		response.Success(writer, userID)
	}
}
