package client

import (
	"net/http"
)

const (
	userAgent      = "Graphael/0.1.0"
	defaultBaseURL = "https://api.graphael.io/"
)

type ClientConfig struct {
	// The base URL for the API. If not provided, the default is used.
	BaseURL *string
	// The client ID to use during authentication.
	ID string
	// The client secret to use during authentication.
	Secret string
}

type Client struct {
	// The base URL for the API.
	baseURL string
	// The HTTP client used to make requests.
	httpClient *http.Client
	// The user agent for the HTTP requests.
	userAgent string
}

// NewClient creates a new Client with the given configuration.
func NewClient(config ClientConfig) *Client {
	var baseURL string

	if config.BaseURL != nil {
		baseURL = *config.BaseURL
	} else {
		baseURL = defaultBaseURL
	}

	if baseURL[len(baseURL)-1] != '/' {
		baseURL += "/"
	}

	return &Client{
		baseURL:    baseURL,
		httpClient: &http.Client{},
		userAgent:  userAgent,
	}
}
