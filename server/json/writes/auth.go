package writes

type AuthSignInResponse struct {
	UserName string `json:"username"`
	Token    string `json:"token"`
}

type AuthSignUpResponse struct {
	UserName string `json:"username"`
	Token    string `json:"token"`
}
