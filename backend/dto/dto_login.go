package dto

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
