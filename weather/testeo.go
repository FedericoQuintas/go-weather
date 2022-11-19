package weather

import (
	"fmt"
    "net/http"
)

type Weather struct {
	city string
	temperature float32
}

func testeoHelper(w http.ResponseWriter, r *http.Request) {
	weather = &Weather{city: "Bologna", temperature: 15.1}
	fmt.Printf("The weather in %s is %f", weather.GetCity, weather.GetTemperature)
}

func (weather Weather) GetCity() string {
	return weather.city
}

func (weather Weather) GetTemperature() float32 {
	return weather.temperature
}
