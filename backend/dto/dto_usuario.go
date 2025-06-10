package dto

type UsuarioCreateDTO struct {
	Nombre   string `json:"nombre" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Rol      string `json:"rol" binding:"required"`
}

type UsuarioResponseDTO struct {
	UsuarioID uint   `json:"usuario_id"`
	Nombre    string `json:"nombre"`
	Email     string `json:"email"`
	Rol       string `json:"rol"`
}

type LoginDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponseDTO struct {
	Token   string             `json:"token"`
	Usuario UsuarioResponseDTO `json:"usuario"`
}
