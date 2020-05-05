package reads

type AuthSignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthSignUpRequest struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
