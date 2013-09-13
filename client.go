package client

import (
	"io"
	"net/http"
)

const (
	rootUri = "https://api.orchestrate.io/v0/"
)

type Client struct {
	HttpClient *http.Client
	AuthToken  string
}

func NewClient(authToken string) *Client {
	httpClient := &http.Client{}

	return &Client{
		HttpClient: httpClient,
		AuthToken:  authToken,
	}
}

func (client Client) doRequest(method, trailingPath string, body io.Reader) (*http.Response, error) {
	req, _ := http.NewRequest(method, rootUri+trailingPath, body)

	req.SetBasicAuth(client.AuthToken, "")

	if method == "PUT" {
		req.Header.Add("Content-Type", "application/json")
	}

	return client.HttpClient.Do(req)
}
