package user

import "github.com/nepp-tumsat/documents-api/model"

type UserGetResponse struct {
	UserName string `json:"username"`
}

func ToUserGetResponse(user model.User) UserGetResponse {
	return UserGetResponse{UserName: user.UserName}
}
