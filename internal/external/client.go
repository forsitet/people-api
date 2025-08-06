package external

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() Client {
	return Client{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

type PersonInfo struct {
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
}

type AgifyResponse struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
}

type GenderizeResponse struct {
	Gender string  `json:"gender"`
	Name   string  `json:"name"`
	Count  int     `json:"count"`
	Prob   float64 `json:"probability"`
}

type NationalizeResponse struct {
	Name    string `json:"name"`
	Country []struct {
		CountryID   string  `json:"country_id"`
		Probability float64 `json:"probability"`
	} `json:"country"`
}

// GetPersonInfo получает информацию о человеке из внешних API
func (c *Client) GetPersonInfo(ctx context.Context, name string) (*PersonInfo, error) {
	info := &PersonInfo{}

	if age, err := c.getAge(ctx, name); err == nil {
		info.Age = age
	}

	if gender, err := c.getGender(ctx, name); err == nil {
		info.Gender = gender
	}

	if nationality, err := c.getNationality(ctx, name); err == nil {
		info.Nationality = nationality
	}

	return info, nil
}

// getAge получает возраст по имени через agify.io
func (c *Client) getAge(ctx context.Context, name string) (int, error) {
	baseURL := "https://api.agify.io"
	params := url.Values{}
	params.Add("name", name)

	req, err := http.NewRequestWithContext(ctx, "GET", baseURL+"?"+params.Encode(), nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to make request: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("Error closing response body: %v", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var agifyResp AgifyResponse
	if err := json.NewDecoder(resp.Body).Decode(&agifyResp); err != nil {
		return 0, fmt.Errorf("failed to decode response: %w", err)
	}

	return agifyResp.Age, nil
}

// getGender получает пол по имени через genderize.io
func (c *Client) getGender(ctx context.Context, name string) (string, error) {
	baseURL := "https://api.genderize.io"
	params := url.Values{}
	params.Add("name", name)

	req, err := http.NewRequestWithContext(ctx, "GET", baseURL+"?"+params.Encode(), nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("Error closing response body: %v", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var genderizeResp GenderizeResponse
	if err := json.NewDecoder(resp.Body).Decode(&genderizeResp); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	switch genderizeResp.Gender {
	case "male":
		return "Мужской", nil
	case "female":
		return "Женский", nil
	default:
		return "", fmt.Errorf("unknown gender: %s", genderizeResp.Gender)
	}
}

// getNationality получает национальность по имени через nationalize.io
func (c *Client) getNationality(ctx context.Context, name string) (string, error) {
	baseURL := "https://api.nationalize.io"
	params := url.Values{}
	params.Add("name", name)

	req, err := http.NewRequestWithContext(ctx, "GET", baseURL+"?"+params.Encode(), nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("Error closing response body: %v", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var nationalizeResp NationalizeResponse
	if err := json.NewDecoder(resp.Body).Decode(&nationalizeResp); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	if len(nationalizeResp.Country) > 0 {
		return nationalizeResp.Country[0].CountryID, nil
	}

	return "", fmt.Errorf("no nationality data available")
}
