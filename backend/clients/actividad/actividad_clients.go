package actividad

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/volpi-figueroa-diazotanez-castellani/proyecto2025/backend/internal/domain"
)

type actividadClient struct {
	client  *http.Client
	baseURL string
}

func NewActividadClient(baseURL string) *actividadClient {
	return &actividadClient{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL: baseURL,
	}
}

func (c *actividadClient) GetActividades(ctx context.Context) ([]domain.Actividad, error) {
	url := fmt.Sprintf("%s/actividades", c.baseURL)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var actividades []domain.Actividad
	if err := json.NewDecoder(resp.Body).Decode(&actividades); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return actividades, nil
}

func (c *actividadClient) GetActividadByID(ctx context.Context, id int) (*domain.Actividad, error) {
	url := fmt.Sprintf("%s/actividades/%d", c.baseURL, id)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var actividad domain.Actividad
	if err := json.NewDecoder(resp.Body).Decode(&actividad); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &actividad, nil
}
