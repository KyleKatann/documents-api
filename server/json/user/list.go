package user

import "github.com/nepp-tumsat/documents-api/model"

type UserListResponse struct {
	Users []UserGetResponse `json:"users"`
}

func ToUserListResponse(users []model.User) UserListResponse {
	var jsonUsers []UserGetResponse
	var jsonUser UserGetResponse
	for _, user := range users {
		jsonUser = ToUserGetResponse(user)
		jsonUsers = append(jsonUsers, jsonUser)
	}
	return UserListResponse{Users: jsonUsers}
}
