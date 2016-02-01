package delighted

import (
	"fmt"

	"github.com/google/go-querystring/query"
)

// SurveyResponses endpoint in delighted, allows you to set query options which is transformed into a query string.
type SurveyResponses struct {
	PerPage      int    `url:"per_page,omitempty"`
	Page         int    `url:"page,omitempty"`
	Since        int    `url:"since,omitempty"`
	Until        int    `url:"until,omitempty"`
	UpdatedSince int    `url:"updated_since,omitempty"`
	UpdatedUntil int    `url:"updated_until,omitempty"`
	Trend        string `url:"trend,omitempty"`
	PersonID     string `url:"person_id,omitempty"`
	PersonEmail  string `url:"person_email,omitempty"` // find response(s) by email
}

// Survey single representation of a survery comment in responses.
type Survey struct {
	ID               string            `json:"id"`
	PersonID         int               `json:"person"`
	Score            int               `json:"score"`
	Comment          string            `json:"comment,omitempty"`
	Permalink        string            `json:"permalink"`
	CreatedAt        int               `json:"created_at"`
	UpdatedAt        int               `json:"updated_at"`
	PersonProperties map[string]string `json:"person_properties,omitemtpy"`
	Notes            []surveyNotes     `json:"notes"`
	Tags             []surveyTags      `json:"tags"`
}

// surveyNotes struct
type surveyNotes struct {
	Notes map[string]string
}

// surveyTags struct
type surveyTags struct {
	Tags map[string]string
}

// SurveyService struct
type SurveyService struct {
	client *Client
}

// GetAll survey responses. Default limit is set to 20 per page.
func (s *SurveyService) GetAll(sr *SurveyResponses) (*[]Survey, error) {
	surveyCollection := []Survey{}

	queryString, err := query.Values(sr)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s%s", "survey_responses.json?", queryString.Encode())

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	_, err = s.client.Do(req, &surveyCollection)

	if err != nil {
		return nil, err
	}

	return &surveyCollection, nil
}
