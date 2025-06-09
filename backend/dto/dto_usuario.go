package dto

type UsuarioCreateDTO struct {
	Nombre   string `json:"nombre" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Rol      string `json:"rol" binding:"required"`
}

type UsuarioResponseDTO struct {
	UsuarioID uint   `json:"usuario_id"`
	Nombre    string `json:"nombre"`
	Email     string `json:"email"`
	Rol       string `json:"rol"`
}

type TokenDto struct {
	Token     string `json:"token"`
	UsuarioID uint   `json:"usuario_id"`
	Rol       string `json:"rol"`
}
