package delighted

import (
	"fmt"

	"github.com/google/go-querystring/query"
)

// Metrics endpoint query options
type Metrics struct {
	Since int    `url:"since,omitempty"`
	Until int    `url:"until,omitempty"`
	Trend string `url:"trend,omitempty"`
}

// MetricResponse payload from delighted
type MetricResponse struct {
	NPS               int `json:"nps,omitempty"`
	PromoterCount     int `json:"promoter_count,omitempty"`
	PromoterPercent   int `json:"promoter_percent,omitemty"`
	PassiveCount      int `json:"passive_count,omitempty"`
	PassivePercent    int `json:"passive_percent,omitempty"`
	DectractorCount   int `json:"detractor_count,omitempty"`
	DectractorPercent int `json:"detractor_percent,omitempty"`
	ResponseCount     int `json:"response_count,omitempty"`
}

// MetricService struct
type MetricService struct {
	client *Client
}

// Get metrics from delighted.
func (s *MetricService) Get(m *Metrics) (*MetricResponse, error) {
	mr := MetricResponse{}

	queryString, err := query.Values(m)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s%s", "metrics.json", queryString.Encode())

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	_, err = s.client.Do(req, &mr)

	if err != nil {
		return nil, err
	}

	return &mr, nil
}
