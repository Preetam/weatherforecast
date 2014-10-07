package weatherforecast

import (
	"bytes"
	"testing"
)

func TestResponse(t *testing.T) {
	r := bytes.NewReader(testApiResponse)
	forecast, err := decodeResponse(r)
	if err != nil {
		t.Fatal(err)
	}

	// check struct
	if len(forecast.List) != 37 {
		t.Errorf("Expected %v in list, got %v", 37, len(forecast.List))
	}

	if forecast.List[7].Rain.ThreeHour != 2 {
		t.Errorf("Expected to see non-zero rain value at index %v", 7)
	}
}
