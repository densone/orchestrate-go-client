package orchestrate

import (
	"encoding/json"
	"fmt"
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

type OrchestrateError struct {
	Status  string
	Message string `json:"message"`
	Locator string `json:"locator"`
}

func NewClient(authToken string) *Client {
	httpClient := &http.Client{}

	return &Client{
		HttpClient: httpClient,
		AuthToken:  authToken,
	}
}

func newError(resp *http.Response) error {
	decoder := json.NewDecoder(resp.Body)
	orchestrateError := new(OrchestrateError)
	decoder.Decode(orchestrateError)

	orchestrateError.Status = resp.Status

	return orchestrateError
}

func (e *OrchestrateError) Error() string {
	return fmt.Sprintf(`%v: %v`, e.Status, e.Message)
}

func (client Client) doRequest(method, trailingPath string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, rootUri+trailingPath, body)

	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(client.AuthToken, "")

	if method == "PUT" {
		req.Header.Add("Content-Type", "application/json")
	}

	return client.HttpClient.Do(req)
}
