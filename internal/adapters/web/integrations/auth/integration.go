package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"yego/internal/platform/config"
)

type Integration interface {
	GetUserEmail(userID string, token string) (string, error)
	GetUserIDByUsername(username string, token string) (string, error)
}

type integration struct {
	baseURL string
	client  *http.Client
}

type UserMeResponse struct {
	Data struct {
		ID    string `json:"id"`
		Email string `json:"email"`
	} `json:"data"`
}

type UserByUsernameResponse struct {
	Data struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	} `json:"data"`
}

func NewIntegration(cfg *config.ConfigurationService) Integration {
	baseURL := cfg.AuthAPIURL
	if baseURL == "" {
		baseURL = "http://localhost:8082"
	}

	return &integration{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

func (i *integration) GetUserEmail(userID string, token string) (string, error) {
	url := fmt.Sprintf("%s/user/me", i.baseURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")

	resp, err := i.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("auth service error: %s", string(body))
	}

	var userResponse UserMeResponse
	if err := json.NewDecoder(resp.Body).Decode(&userResponse); err != nil {
		return "", err
	}

	return userResponse.Data.Email, nil
}

func (i *integration) GetUserIDByUsername(username string, token string) (string, error) {
	url := fmt.Sprintf("%s/user/%s", i.baseURL, username)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")

	resp, err := i.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("auth service error: %s", string(body))
	}

	var userResponse UserByUsernameResponse
	if err := json.NewDecoder(resp.Body).Decode(&userResponse); err != nil {
		return "", err
	}

	return userResponse.Data.ID, nil
}
