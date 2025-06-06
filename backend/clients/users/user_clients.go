package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"backend/models"
)

type UserClient struct {
	baseURL string
	client  *http.Client
}

type UserClientInterface interface {
	CreateUser(user models.Usuario) error
	GetUserByEmail(email string) (models.Usuario, error)
	GetUserByID(id uint) (models.Usuario, error)
	UpdateUser(user models.Usuario) error
	DeleteUser(id uint) error
}

var (
	DefaultUserClient UserClientInterface
)

func init() {
	DefaultUserClient = &UserClient{
		baseURL: "http://your-api-base-url", // Cambia esto por tu URL base
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *UserClient) CreateUser(user models.Usuario) error {
	body, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("error marshaling user: %w", err)
	}

	resp, err := c.client.Post(fmt.Sprintf("%s/usuarios", c.baseURL), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func (c *UserClient) GetUserByEmail(email string) (models.Usuario, error) {
	resp, err := c.client.Get(fmt.Sprintf("%s/usuarios/email/%s", c.baseURL, email))
	if err != nil {
		return models.Usuario{}, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return models.Usuario{}, nil
	}

	if resp.StatusCode != http.StatusOK {
		return models.Usuario{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var user models.Usuario
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return models.Usuario{}, fmt.Errorf("error decoding response: %w", err)
	}

	return user, nil
}

func (c *UserClient) GetUserByID(id uint) (models.Usuario, error) {
	resp, err := c.client.Get(fmt.Sprintf("%s/usuarios/%d", c.baseURL, id))
	if err != nil {
		return models.Usuario{}, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return models.Usuario{}, nil
	}

	if resp.StatusCode != http.StatusOK {
		return models.Usuario{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var user models.Usuario
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return models.Usuario{}, fmt.Errorf("error decoding response: %w", err)
	}

	return user, nil
}

func (c *UserClient) UpdateUser(user models.Usuario) error {
	body, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("error marshaling user: %w", err)
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/usuarios/%d", c.baseURL, user.ID), bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func (c *UserClient) DeleteUser(id uint) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/usuarios/%d", c.baseURL, id), nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
