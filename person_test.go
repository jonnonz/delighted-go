package delighted

import (
	"errors"
	"testing"
)

var (
	mockCreatePersonResponse string = `{"id": "1","email": "jony@appleseed.com","name": "John Appleseed","survey_scheduled_at": 1454321359}`
)

func TestCreate(t *testing.T) {
	_, c := mockServer(200, mockCreatePersonResponse)
	p := Person{Name: "John Appleseed", Email: "jony@appleseed.com"}
	n, err := c.PersonService.Create(&p)
	if err != nil {
		t.Errorf("Test failed, error creating person.", err)
	}

	if n.Email != p.Email {
		t.Errorf("Test failed, new person email did not match expected person email.")
	}
}

func TestCreateNoEmail(t *testing.T) {
	_, c := mockServer(422, mockUnprossessableError)
	p := Person{Name: "John Appleseed"}
	_, err := c.PersonService.Create(&p)

	expectedError := errors.New("Email address is required when creating a person.")

	if err.Error() != expectedError.Error() {
		t.Errorf("Test failed, blank email allowed when creating a person.")
	}
}

func TestUnsubscribe(t *testing.T) {
	_, c := mockServer(200, mockSuccessResponse)
	p := Person{Email: "jony@appleseed.com"}
	r, err := c.PersonService.Unsubscribe(&p)
	if err != nil {
		t.Errorf("Test failed, error unsubscribing person.", err)
	}

	if r != true {
		t.Errorf("Test failed, unsubscribing a person did not return true.")
	}
}

func TestUnSubcribeNoEmail(t *testing.T) {
	_, c := mockServer(200, mockCreatePersonResponse)
	p := Person{Name: "John Appleseed"}
	_, err := c.PersonService.Unsubscribe(&p)

	expectedError := errors.New("Email address is required when ubsubscribing a person.")

	if err.Error() != expectedError.Error() {
		t.Errorf("Test failed, blank email allowed when unsubscribing a person.")
	}
}

func TestRemovePendingSurvey(t *testing.T) {
	_, c := mockServer(200, mockSuccessResponse)
	p := Person{Email: "jony@appleseed.com"}
	r, err := c.PersonService.RemovePendingSurvey(&p)
	if err != nil {
		t.Errorf("Test failed, error when removing a pending survey for a person.", err)
	}

	if r != true {
		t.Errorf("Test failed, removing a pending survey for a person did not return true.")
	}
}
