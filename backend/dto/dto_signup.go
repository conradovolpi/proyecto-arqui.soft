package dto

type SignUpRequest struct {
	Nombre   string `json:"nombre"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	Message string `json:"message"`
}
