package main

import (
	"log"
	"net"

	pb "github.com/jasonlarson44/weather-service/protobuf"
	"github.com/jasonlarson44/weather-service/src/util"
	"github.com/jasonlarson44/weather-service/src/weather"
	"google.golang.org/grpc"
)

func main() {
	config, err := util.LoadConfig(".")
	port := ":" + config.Port

	log.Printf("starting server on port %s", port)
	listener, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterWeatherServiceServer(s, &weather.Server{Config: config, UnimplementedWeatherServiceServer: pb.UnimplementedWeatherServiceServer{}})
	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
