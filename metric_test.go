package delighted

import (
	"fmt"

	"testing"
)

var mockMetricResponse = `{"nps":100,"promoter_count":2,"promoter_percent":100.0,"passive_count":0,"passive_percent":0.0,"detractor_count":0,"detractor_percent":0.0,"response_count":2}`

// TestMetricResponse Error and Message
func TestMetricResponse(t *testing.T) {
	_, client := mockServer(200, mockMetricResponse)

	m := Metrics{}
	d, err := client.MetricService.Get(&m)
	if err != nil {
		fmt.Println(err.Error())
		t.Errorf("Test failed, unable to get response from mock client.", err)
	}

	if d.NPS != 100 || d.PromoterCount != 2 {
		t.Errorf("Test failed, mock response did not match the expected values.")
	}
}
