package user

import "github.com/nepp-tumsat/documents-api/model"

type UserListResponse struct {
	Users []UserGetResponse `json:"users"`
}

func toJsonUserList(users []model.User) []UserGetResponse {
	var JsonUsers []UserGetResponse
	for _, user := range users {
		user = toJsonUser(user)
		users = append(users, user)
	}
	return UserListResponse{Users: JsonUsers}
}
