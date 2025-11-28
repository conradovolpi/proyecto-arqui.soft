package services

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"os"

	"backend/clients/usuario"
	"backend/dto"
	"backend/models"
	"backend/utils"
)

type UsuarioService interface {
	Create(dto.UsuarioCreateDTO) (*dto.UsuarioResponseDTO, utils.ApiError)
	Login(dto.LoginDTO) (*dto.LoginResponseDTO, utils.ApiError)
	GetByID(uint) (*dto.UsuarioResponseDTO, utils.ApiError)
	GetAll() ([]dto.UsuarioResponseDTO, utils.ApiError)
}

type usuarioService struct {
	client    usuario.UsuarioClientInterface
	jwtSecret string
}

func NewUsuarioService() UsuarioService {
	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		panic("JWT_SECRET_KEY no definida en entorno")
	}
	return &usuarioService{
		client:    usuario.UsuarioClient,
		jwtSecret: secret,
	}
}

func (s *usuarioService) Create(u dto.UsuarioCreateDTO) (*dto.UsuarioResponseDTO, utils.ApiError) {
	// Hashear la contraseña con SHA256
	hash := sha256.Sum256([]byte(u.Password))
	hashedPassword := hex.EncodeToString(hash[:])

	log.Printf("Create: Creando usuario - Email: %s, Hash length: %d, Hash: %s", u.Email, len(hashedPassword), hashedPassword)

	usuario := &models.Usuario{
		Nombre:   u.Nombre,
		Email:    u.Email,
		Password: hashedPassword,
		Rol:      u.Rol,
	}

	if err := s.client.CreateUser(usuario); err != nil {
		log.Printf("Create: Error al crear usuario: %v", err)
		return nil, utils.NewInternalServerApiError("Error creando usuario", err)
	}

	// Verificar que el usuario se guardó correctamente leyéndolo de nuevo
	usuarioVerificado, err := s.client.GetByID(usuario.UsuarioID)
	if err != nil {
		log.Printf("Create: Advertencia - No se pudo verificar el usuario creado: %v", err)
	} else {
		log.Printf("Create: Usuario verificado - ID: %d, Email: %s, Hash guardado length: %d", 
			usuarioVerificado.UsuarioID, usuarioVerificado.Email, len(usuarioVerificado.Password))
		if usuarioVerificado.Password != hashedPassword {
			log.Printf("Create: ERROR - El hash guardado no coincide con el hash calculado!")
			log.Printf("Create: Hash calculado: %s", hashedPassword)
			log.Printf("Create: Hash guardado: %s", usuarioVerificado.Password)
		}
	}

	log.Printf("Create: Usuario creado exitosamente - ID: %d, Email: %s", usuario.UsuarioID, usuario.Email)

	return &dto.UsuarioResponseDTO{
		UsuarioID: usuario.UsuarioID,
		Nombre:    usuario.Nombre,
		Email:     usuario.Email,
		Rol:       usuario.Rol,
	}, nil
}

func (s *usuarioService) Login(loginDTO dto.LoginDTO) (*dto.LoginResponseDTO, utils.ApiError) {
	log.Printf("Login: Buscando usuario con email: %s", loginDTO.Email)
	usuario, err := s.client.GetByEmail(loginDTO.Email)
	if err != nil {
		log.Printf("Login: Error al buscar usuario: %v", err)
		return nil, utils.NewUnauthorizedApiError("Email o contraseña incorrectos")
	}

	log.Printf("Login: Usuario encontrado - ID: %d, Email: %s", usuario.UsuarioID, usuario.Email)
	log.Printf("Login: Hash almacenado en BD - Length: %d, Hash: %s", len(usuario.Password), usuario.Password)

	// Hashear la contraseña proporcionada con SHA256 para comparar
	hash := sha256.Sum256([]byte(loginDTO.Password))
	hashedPassword := hex.EncodeToString(hash[:])

	log.Printf("Login: Hash calculado de password ingresado - Length: %d, Hash: %s", len(hashedPassword), hashedPassword)
	log.Printf("Login: Comparando contraseñas hasheadas...")
	log.Printf("Login: Hash BD == Hash calculado? %v", usuario.Password == hashedPassword)

	// Limpiar espacios en blanco por si acaso
	storedHash := usuario.Password
	calculatedHash := hashedPassword

	if storedHash != calculatedHash {
		log.Printf("Login: ERROR - Los hashes no coinciden")
		log.Printf("Login: Hash almacenado: '%s'", storedHash)
		log.Printf("Login: Hash calculado: '%s'", calculatedHash)
		return nil, utils.NewUnauthorizedApiError("Email o contraseña incorrectos")
	}

	log.Printf("Login: Contraseña correcta, generando token...")
	token := utils.GenerateJWT(usuario.UsuarioID, usuario.Rol, s.jwtSecret)

	log.Printf("Login: Token generado exitosamente para usuario ID: %d", usuario.UsuarioID)
	return &dto.LoginResponseDTO{
		Token: token,
		Usuario: dto.UsuarioResponseDTO{
			UsuarioID: usuario.UsuarioID,
			Nombre:    usuario.Nombre,
			Email:     usuario.Email,
			Rol:       usuario.Rol,
		},
	}, nil
}

func (s *usuarioService) GetByID(id uint) (*dto.UsuarioResponseDTO, utils.ApiError) {
	usuario, err := s.client.GetByID(id)
	if err != nil {
		return nil, utils.NewNotFoundApiError("Usuario no encontrado")
	}

	return &dto.UsuarioResponseDTO{
		UsuarioID: usuario.UsuarioID,
		Nombre:    usuario.Nombre,
		Email:     usuario.Email,
		Rol:       usuario.Rol,
	}, nil
}

func (s *usuarioService) GetAll() ([]dto.UsuarioResponseDTO, utils.ApiError) {
	usuarios, err := s.client.GetAll()
	if err != nil {
		return nil, utils.NewInternalServerApiError("Error recuperando usuarios", err)
	}

	var dtos []dto.UsuarioResponseDTO
	for _, u := range usuarios {
		dtos = append(dtos, dto.UsuarioResponseDTO{
			UsuarioID: u.UsuarioID,
			Nombre:    u.Nombre,
			Email:     u.Email,
			Rol:       u.Rol,
		})
	}
	return dtos, nil
}
