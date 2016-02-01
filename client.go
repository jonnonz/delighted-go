package delighted

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	// LibVersion used to be sent with User-Agent
	LibVersion = "0.2beta"
)

var (
	// DelightedAPIKey for delighted.
	DelightedAPIKey string
	BaseURL         string = "https://api.delighted.com/v1/%s"
)

// Client handles communication between lib and delighted api.
type Client struct {
	client    *http.Client
	UserAgent string

	PersonService *PersonService
	SurveyService *SurveyService
	MetricService *MetricService
}

// NewClient returns a new client.
func NewClient(APIKey string, httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	if APIKey == "" {
		return nil, errors.New("No API key has been provided.")
	}

	DelightedAPIKey = APIKey

	userAgent := fmt.Sprintf("Go-Delighted-Version-%s", LibVersion)

	c := Client{client: httpClient, UserAgent: userAgent}

	c.PersonService = &PersonService{&c}
	c.SurveyService = &SurveyService{&c}
	c.MetricService = &MetricService{&c}

	return &c, nil
}

// NewRequest to delighted api.
func (c *Client) NewRequest(method, uri string, body interface{}) (*http.Request, error) {
	url := fmt.Sprintf(BaseURL, uri)
	buf := &bytes.Buffer{}
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(DelightedAPIKey, "")
	req.Header.Add("User-Agent", c.UserAgent)
	req.Header.Add("Content-Type", "application/json")

	return req, nil
}

// Do request
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = HandleResponse(resp)
	if err != nil {
		return resp, err
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	return resp, nil
}

// HandleResponse from delighted app.
func HandleResponse(r *http.Response) error {
	switch r.StatusCode {
	case 201, 200, 202:
		return nil
	case 401:
		return errors.New("Authentication error.")
	case 406:
		return errors.New("Unsupported request format.")
	case 422:
		return errors.New("Resource validation error.")
	case 503:
		return errors.New("Unable to communicate with server.")
	default:
		return errors.New("Unknown api error.")
	}
}
