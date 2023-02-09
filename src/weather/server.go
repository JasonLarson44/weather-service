package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	pb "github.com/jasonlarson44/weather-service/protobuf"
	"github.com/jasonlarson44/weather-service/src/util"
)

type Server struct {
	Config util.Config
	pb.UnimplementedWeatherServiceServer
}

// Type to map OpenWeather OneApi json response to Go types
type OWResp struct {
	Lat     float64            `json:"lat"`
	Lon     float64            `json:"lon"`
	Current OWRespCurrent      `json:"current"`
	Alerts  []*pb.WeatherAlert `json:"alerts"`
}

// Type for the `current` model of the OpenWeather OneApi JSON response
type OWRespCurrent struct {
	Temp    float64         `json:"temp"`
	Weather []OWRespWeather `json:"weather"`
}

// Type for the `current.weather` model of the OpenWeather OneApi JSON response
type OWRespWeather struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
}

// Implements the GetWeather api endpoint. Expects float64 latitude and longitude and an optional units string.
// Returns Current weather conditions at given location
func (s *Server) GetWeather(ctx context.Context, in *pb.WeatherRequest) (*pb.WeatherResponse, error) {
	log.Printf("Received getWeather request for latitude: %f, longitude: %f", in.Lat, in.Lon)
	urlParams := fmt.Sprintf("?lat=%f&lon=%f&exclude=%s&units=%s&appid=%s", in.Lat, in.Lon, "hourly,daily,minutely", in.GetUnits().String(), s.Config.ApiKey)
	res, err := http.Get(s.Config.OWBaseUri + urlParams)
	if err != nil {
		log.Printf("Failed to get weather data with error %v", err)
	}

	if res.StatusCode != 200 {
		log.Printf("Failed to get weather data with status code: %d\n", res.StatusCode)
		return &pb.WeatherResponse{}, fmt.Errorf("Failed to get weather data from OpenWeather with code: %d", res.StatusCode)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}

	var data OWResp
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Printf("Failed to unmarshal: %+v", err)
	}

	return &pb.WeatherResponse{
		Conditions: data.Current.Weather[0].Description,
		Temp:       tempToRating(data.Current.Temp, in.GetUnits().String()),
		Alerts:     data.Alerts,
	}, nil
}

// Takes the temp value and unit of measurement to produce a temperature rating of cold, moderate or hot
func tempToRating(temp float64, units string) string {
	switch strings.ToLower(units) {
	case "metric": //celsius
		switch {
		// <= 50F
		case temp <= 10.00:
			return "Cold"
			// between 50F and 75F
		case temp > 10.00 && temp < 23.89:
			return "Moderate"
			// > 75F
		case temp >= 23.89:
			return "Hot"
		default:
			return "Cold"
		}
	case "standard": //kelvin
		switch {
		// <= 50F
		case temp <= 283.15:
			return "Cold"
			// between 50F and 75F
		case temp > 283.15 && temp < 297.04:
			return "Moderate"
			// > 75F
		case temp >= 297.04:
			return "Hot"
		default:
			return "Cold"
		}
	default: // Default to fahrenheit
		switch {
		case temp <= 50.0:
			return "Cold"
		case temp > 50.0 && temp < 75.0:
			return "Moderate"
		case temp >= 75.0:
			return "Hot"
		default:
			return "Cold"
		}

	}
}
