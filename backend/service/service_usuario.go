package service

import (
	"backend/clients"
	"backend/dao"
	"backend/dto"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UsuarioService struct{}

type UsuarioServiceInterface interface {
	Signup(signUp dto.SignUpRequest) error
	Login(User dto.UserLoginRequest) dto.UserLoginResponse
}

var (
	UsuarioServiceInterfaceInstance UsuarioServiceInterface
)

func init() {
	UsuarioServiceInterfaceInstance = &UsuarioService{}
}

// Métodos de signup
func (s *UsuarioService) Signup(signUp dto.SignUpRequest) error {
	// Verificar si el usuario ya existe
	_, err := clients.ObtainUserByEmail(signUp.Email)
	if err == nil {
		return errors.New("user already exists")
	}

	// Hash de la contraseña
	hashedPassword, err := hashPassword(signUp.Password)
	if err != nil {
		return err
	}

	// Crear un nuevo usuario
	newUser := &dao.Usuario{
		Email:    signUp.Email,
		Password: hashedPassword,
		Nombre:   signUp.Nombre,
		Rol:      "cliente", // Asignar rol por defecto
	}

	// Guardar el usuario en la base de datos
	if err := clients.CreateUser(newUser); err != nil {
		return err
	}

	return nil
}

// Métodos de login
func (s *UsuarioService) Login(User dto.UserLoginRequest) (response dto.UserLoginResponse) {
	client := dao.Usuario{
		Email:    User.Email,
		Password: User.Password,
	}

	// Verificar si el usuario existe
	userDAO, err := clients.ObtainUserByEmail(User.Email)
	if err != nil {
		response.Message = "Invalid email or password"
		return response
	}

	// Comparar la contraseña enviada con la contraseña guardada en la base de datos
	err = bcrypt.CompareHashAndPassword([]byte(userDAO.Password), []byte(client.Password))
	if err != nil {
		response.Message = "Invalid email or password"
		return response
	}

	// Generar un token JWT
	tokenString, err := generateJWT(userDAO.Email, userDAO.ID)
	if err != nil {
		response.Message = "Error generating token"
		return response
	}

	response.Message = "Login successful"
	response.Token = tokenString
	return response
}

// Funcion para encriptar la contraseña
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func generateJWT(email string, userId uint) (string, error) {
	// Generar un token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":    email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
