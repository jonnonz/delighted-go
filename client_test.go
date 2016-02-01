package delighted

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"

	"testing"
)

var (
	mockAuthError           string = `{"status":401,"message":"Unauthorized"}`
	mockUnprossessableError string = `{"status": 422,"message": "Unprocessable Entity"}`
	mockSuccessResponse     string = `{"status": "ok"}`
)

// mockServer for testing.
func mockServer(code int, body string) (*httptest.Server, *Client) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, body)
	}))

	BaseURL = fmt.Sprintf("%s/%%s", server.URL)

	tr := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}
	httpClient := &http.Client{Transport: tr}

	client, _ := NewClient("test-token", httpClient)

	return server, client
}

// TestNewClientNoAPIKey Error and Message
func TestNewClientNoAPIKey(t *testing.T) {

	_, err := NewClient("", nil)

	expectedError := errors.New("No API key has been provided.")

	if err.Error() != expectedError.Error() {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", err, expectedError)
	}
}

func TestClientWithAPIKey(t *testing.T) {
	transport := &http.Transport{}

	// Make a http.Client with the transport
	httpClient := &http.Client{Transport: transport}

	c, err := NewClient("12345678", httpClient)

	if err != nil {
		t.Errorf("Test failed, count not create client with correct variables provided.")
	}

	expectedClient := Client{}

	if reflect.TypeOf(c) != reflect.TypeOf(&expectedClient) {
		t.Errorf("Test failed, client provided was not correct client type.")
	}
}

func TestUnAuthenticationRequestError(t *testing.T) {
	server, c := mockServer(401, mockAuthError)
	defer server.Close()

	expectedError := errors.New("Authentication error.")

	mockPerson := Person{Email: "test@test.com"}
	_, err := c.PersonService.Create(&mockPerson)

	if expectedError.Error() != err.Error() {
		t.Errorf("Test failed, 401 unauthenticated test failed.")
	}
}

func TestUnProcessableRequestError(t *testing.T) {
	server, c := mockServer(422, mockUnprossessableError)
	defer server.Close()

	expectedError := errors.New("Resource validation error.")

	mockPerson := Person{Email: "test@test.com"}
	_, err := c.PersonService.Create(&mockPerson)

	if expectedError.Error() != err.Error() {
		t.Errorf("Test failed, 422 unprocessable failed.")
	}
}

func TestSuccessfulRequest(t *testing.T) {
	server, c := mockServer(200, mockSuccessResponse)
	defer server.Close()

	mockPerson := Person{Email: "test@test.com"}
	_, err := c.PersonService.Create(&mockPerson)

	if err != nil {
		t.Errorf("Test failed, 200 success response failed.")
	}
}
