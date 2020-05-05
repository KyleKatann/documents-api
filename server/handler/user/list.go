package user

import (
	"net/http"

	"github.com/nepp-tumsat/documents-api/infrastructure"
	"github.com/nepp-tumsat/documents-api/infrastructure/persistence"
	"github.com/nepp-tumsat/documents-api/model"
	"github.com/nepp-tumsat/documents-api/server/response"
	"golang.org/x/xerrors"
)

func HandleUserList() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		userRepo := persistence.NewUserDB(infrastructure.DB)

		users, err := userRepo.SelectAll()

		if err != nil {
			err = xerrors.Errorf("Error in repository: %v", err)
		}

		response.Success(writer, userListResponse{Users: users})
	}
}

type userListResponse struct {
	Users []model.User `json:"users"`
}
