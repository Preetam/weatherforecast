package weatherforecast

import (
	"encoding/json"
	"io"
	"net/http"
)

func FetchForecast(q string) (*WeatherForecast, error) {

	res, err := http.Get("http://api.openweathermap.org/data/2.5/forecast?q=")
	if err != nil {
		return nil, err
	}

	forecast, err := decodeResponse(res.Body)
	res.Body.Close()

	return forecast, err
}

func decodeResponse(response io.Reader) (*WeatherForecast, error) {
	forecast := WeatherForecast{}
	dec := json.NewDecoder(response)
	err := dec.Decode(&forecast)
	return &forecast, err
}
