package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"strconv"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoch    int64   `json:"time_epoch"`
				TempC        float64 `json:"temp_c"`
				ChanceOfRain float64 `json:"chance_of_rain"`
				Condition    struct {
					Text string `json:"text"`
				} `json:"condition"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {
	region := "Gurgaon"
	skips := 6
	days := "1"

	if len(os.Args) >= 2 {
		temp := os.Args[1]
		if temp == "ggn" {
			region = "Gurgaon"
		} else if temp == "del" || temp == "rhni" {
			region = "Rohini"
		} else {
			region = temp
		}
	}

	if len(os.Args) >= 3 {
		var err error
		skips, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Error converting skips to integer: %v", err)
		}
	}

	if len(os.Args) >= 4 {
		days = os.Args[3]
	}

	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Read the WEATHER_API_KEY from environment variables
	apiKey, found := os.LookupEnv("WEATHER_API_KEY")
	if !found {
		log.Fatalf("WEATHER_API_KEY not found in .env")
	}

	url := fmt.Sprintf("https://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=%s&aqi=no&alerts=no", apiKey, region, days)

	res, err := http.Get(url)

	if err != nil {
		log.Fatalf("Error fetching weather data: %v", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("Error reading resonse body: %v", err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	location, current, condition := weather.Location, weather.Current, weather.Current.Condition.Text
	forecast := weather.Forecast.Forecastday

	metaMessage := fmt.Sprintf("Location: %s, %s\nCurrent Temperature: %.0f°C\nCurrent Condition: %s\n", location.Name, location.Country, current.TempC, condition)

	color.Cyan(metaMessage)

	for _, day := range forecast {
		// date in dd mm yy and day
		color.Yellow(fmt.Sprintf("\nDate: %s\n", time.Unix(day.Hour[0].TimeEpoch, 0).Format("Mon, 02 Jan 2006")))
		for index, hour := range day.Hour {
			t := time.Unix(hour.TimeEpoch, 0)

			message := fmt.Sprintf("Time: %s, Temp: %.0f°C, Rain: %.0f, Condition: %s", t.Format("15:04"), hour.TempC, hour.ChanceOfRain, hour.Condition.Text)

			// skip the hours
			if index%int(skips) != 0 {
				continue
			}

			if hour.ChanceOfRain > 10 {
				color.Red(message)
			} else {
				color.Green(message)
			}
		}

	}

}
