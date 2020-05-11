package jokes

import (
	"encoding/json"
	"fmt"
	"main/internal/api"
	"net/http"
)

const getJokePath = "/api?format=json"

// JokeClient is a joke API client
type JokeClient struct {
	url string
}

func NewJokeClient(baseURL string) *JokeClient {
	return &JokeClient{
		url: baseURL,
	}
}

func (jc *JokeClient) GetJoke() (*api.JokeResponse, error) {
	urlPath := jc.url + getJokePath

	resp, err := http.Get(urlPath)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request status: %s", http.StatusText(resp.StatusCode))
	}

	var data api.JokeResponse

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
