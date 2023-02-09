/* Implements a basic client to interact with the gRPC weather server
*
 */

package main

import (
	"context"
	"log"
	"time"

	pb "github.com/jasonlarson44/weather-service/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = "localhost:8080"
	// Configure request below here
	latitude = 29.95
	longitude = -90.07
	units = pb.UnitsType_IMPERIAL
)

func main() {
	// Get connection to server
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewWeatherServiceClient(conn)

	// Make API call
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	req := &pb.WeatherRequest{Lat: latitude, Lon: longitude, Units: &units}

	res, err := c.GetWeather(ctx, req)
	if err != nil {
		log.Fatalf("Failed to get weather with error: %v", err)
	}
	log.Printf("Weather response for %f, %f:\n", latitude, longitude)
	printRes(res)
}

func printRes(wr *pb.WeatherResponse) {
	log.Printf("Temperature: %s\n", wr.Temp)
	log.Printf("Weather Conditions: %s\n\n", wr.Conditions)
	log.Printf("Ongoing Weather Alerts:\n")
	for i, alert := range wr.Alerts{
		log.Printf("Event %d: %s\n", i+1, alert.Event)
		log.Printf("Sender Name: %s\n", alert.SenderName)
		log.Printf("Start time: %d\n", alert.Start)
		log.Printf("End time: %d\n", alert.End)
		log.Printf("Description: %s\n\n", alert.Description)
	}
}

