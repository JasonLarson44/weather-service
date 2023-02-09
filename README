# Go gRPC Weather Server

A Go gRPC server that serves a single endpoint to retrieve the current weather conditions for a given latitude/longitude coordinate.

This project serves a gRPC API and utilizes protobuf. You may find the .proto and protobuf generated go files in /protobuf. All of the server code and tests reside in /src. The /client directory provides a simple go client for making the rpc call to the server's GetWeather endpoint. 

Configuration is handled by [spf13/viper](https://github.com/spf13/viper) and is done in the `app.env` file at the root of the project. The /src/util/config.go file provides type to load the configuration values into objects to be passed around as needed.

## API
* /GetWeather - This endpoint takes in a latitude and longitude coordinate in the form of two float64 values, as well as an optional units field to use imperial, metric or standard units of measure. It returns the current weather data, along with any active weather alerts at the location.

Example Request body:
```
{
    "lat": 21.31,
    "lon": -157.85,
    "units": "IMPERIAL"
}
```

Example Response Body:
```
{
    "alerts": [
        {
            "sender_name": "NWS Honolulu (Hawaii)",
            "event": "Small Craft Advisory",
            "start": "1675950960",
            "end": "1676088000",
            "description": "...SMALL CRAFT ADVISORY REMAINS IN EFFECT UNTIL 6 PM HST FRIDAY...\n* WHAT...East winds 20 to 30 kt. Seas 10 to 15 feet.\n* WHERE...Kauai Northwest Waters, Kauai Windward Waters, Kauai\nLeeward Waters, Kauai Channel, Oahu Windward Waters, Oahu\nLeeward Waters, Maui County Windward Waters, Maui County\nLeeward Waters and Big Island Windward Waters.\n* WHEN...Until 6 PM HST Friday.\n* IMPACTS...Conditions will be hazardous to small craft."
        },
        {
            "sender_name": "NWS Honolulu (Hawaii)",
            "event": "Gale Warning",
            "start": "1675966200",
            "end": "1676003400",
            "description": "Hawaiian offshore waters beyond 40 nautical miles out to 240\nnautical miles including the portion of the Papahanaumokuakea\nMarine National Monument east of French Frigate Shoals\nSeas given as significant wave height, which is the average height\nof the highest 1/3 of the waves. Individual waves may be more than\ntwice the significant wave height.\n...GALE WARNING...\n.THIS AFTERNOON...E winds 25 to 35 kt, highest S waters. Seas 10\nto 15 ft.\n.TONIGHT THROUGH FRIDAY NIGHT...E winds 25 to 35 kt, highest SE.\nSeas 12 to 16 ft. Isolated thunderstorms far SE.\n.SATURDAY THROUGH MONDAY...E winds 15 to 25 kt. Seas 10 to 15 ft\nsubsiding to 8 to 12 ft Sunday."
        }
    ],
    "conditions": "overcast clouds",
    "temp": "Hot"
}
```

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

You need to have the following software installed on your machine:

- Go >= 1.19
- Protocol Buffers >= 3.19.4
- gRPC Go plugin >= 1.53.0

This API makes use of the [OpenWeather One Call API 3.0](https://openweathermap.org/api/one-call-3) to provide weather data. In order for it to properly run you will need an OpenWeather API key with a subscription to the ["One Call by Call"](https://openweathermap.org/price) plan.

Once you have your active OpenWeather API key, you will need to configure you `app.env` file at the root of the repo. There is a template available in `template.env`. Run `cp template.env app.env` to copy the template and start your `app.env` file. Replace the placeholder value for `OW_API_KEY` in `app.env` with your OpenWeather api key. The app.env file has been added to the .gitignore to keep it out of source control, but take care to not commit app.env to source control as it will expose your api key. 

### Installing

Clone this repository to your local machine

Change directories to the root of your newly cloned repository
Run `go get` from the root of the repository to download and install the project's dependencies. 

### Start Server
Once prerequisites have been met and dependencies have been installed the server can be started using the following command from the root of the repository:

`go run ./src/main.go`

The server defaults to listening on port 8080, but that can be changed in `app.env` file.

### Hitting the GetWeather endpoint
Once the server is up and running, you should be able to hit the GetWeather endpoint using any tool that supports gRPC([Postman](https://www.postman.com/), [BloomRPC](https://github.com/bloomrpc/bloomrpc)). I have also included a example Go client in this repo for testing. It can be found in `/client/main.go `. It is a simple program that just makes a single request to the GetWeather endpoint. The port and request parameters can be configured in the var section near the top of the file. In order to use the client to make a call to the server., simply run: `go run ./client/main.go` from the root of the repo while the server is up and running. It should hit the server with your request and print out the response in your terminal. Ensure that the port the server is running on matches the port variable in `/client/main.go`

### Testing
I have added some very basic tests for the weather module. They can be run from the root of the project with the following command:
`go test ./src/...`

### Formatting
This repo makes use of the `go fmt` command. Before making any commits please run `go fmt ./src/...` from the root

