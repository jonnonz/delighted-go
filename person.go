package delighted

import (
	"errors"
	"fmt"
)

// Person representation in golang.
type Person struct {
	ID         string            `json:"id,omitempty"`
	Email      string            `json:"email"`
	Name       string            `json:"name,omitempty"`
	Delay      int               `json:"delay,omitempty"`
	Properties map[string]string `json:"properties,omitempty"`
	Send       bool              `json:"send,omitempty"`
	LastSentAt int               `json:"last_sent_at,omitempty"`
}

// PersonSurveyResponse from action against a person's survey
type PersonSurveyResponse struct {
	Ok string `json:"ok,omitempty"`
}

// PersonUnsubscribePayload for unsubscribing a person.
type PersonUnsubscribePayload struct {
	Email string `json:"person_email"`
}

// PeopleService struct
type PersonService struct {
	client *Client
}

// Create a new person for a survey.
func (s *PersonService) Create(p *Person) (*Person, error) {
	if p.Email == "" {
		return p, errors.New("Email address is required when creating a person.")
	}
	newPerson := Person{}

	req, err := s.client.NewRequest("POST", "people.json", p)
	if err != nil {
		return p, err
	}

	_, err = s.client.Do(req, &newPerson)
	if err != nil {
		return &newPerson, err
	}

	return &newPerson, nil
}

// Unsubscribe a person from a survey, returns true of false if person was sucessfully Unsubscribed
func (s *PersonService) Unsubscribe(p *Person) (bool, error) {
	if p.Email == "" {
		return false, errors.New("Email address is required when ubsubscribing a person.")
	}

	PersonUnsubscribePayload := PersonUnsubscribePayload{Email: p.Email}

	req, err := s.client.NewRequest("POST", "unsubscribes.json", PersonUnsubscribePayload)
	if err != nil {
		return false, err
	}

	_, err = s.client.Do(req, nil)
	if err != nil {
		return false, err
	}

	return true, nil
}

// RemovePendingSurvey from a person who already has a pending survey
func (s *PersonService) RemovePendingSurvey(p *Person) (bool, error) {
	if p.Email == "" {
		return false, errors.New("Email address is required when unsubscribing a person.")
	}

	psr := PersonSurveyResponse{}
	url := fmt.Sprintf("people/%s/survey_requests/pending.json", p.Email)

	req, err := s.client.NewRequest("DELETE", url, nil)
	if err != nil {
		return false, err
	}

	_, err = s.client.Do(req, &psr)
	if err != nil {
		return false, err
	}

	return true, nil
}
