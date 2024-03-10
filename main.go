package main

import (
	"context"
	"fmt"
	"goEpexSpot/internal/clients/rte/swagger"
	"goEpexSpot/internal/clients/rte/token"
	"log"
	"net/http"
	"os"
)

func main() {

	basicAuthUsername := os.Getenv("BASIC_AUTH_USERNAME")
	basicAuthPassword := os.Getenv("BASIC_AUTH_PASSWORD")

	if basicAuthUsername == "" || basicAuthPassword == "" {
		log.Fatal("basic auth credentials are not set in .env file. Please set BASIC_AUTH_USERNAME and BASIC_AUTH_PASSWORD.")
	}

	bearerToken, err := token.FetchBearerToken(basicAuthUsername, basicAuthPassword)
	if err != nil {
		log.Fatalf("Error while fetching bearer token: %s\n", err)
	}

	// Create a new API client configuration
	cfg := swagger.NewConfiguration()
	cfg.HTTPClient = &http.Client{}
	cfg.AddDefaultHeader("Authorization", "Bearer "+bearerToken) // Add the Authorization header with the Bearer token

	// Create a new API client with the configuration
	client := swagger.NewAPIClient(cfg)

	// Assuming there's a GetMarketData method on your DefaultApi service
	// Adjust the method and parameters according to your actual API
	data, resp, err := client.DefaultApi.GetFrancePowerExchanges(context.Background())
	_ = resp
	// fmt.Printf("Response: %+v\n", resp)
	if err != nil {
		log.Printf("Error while calling API: %+v", err)
	} else {
		fmt.Printf("Market Data: %+v\n", data)
	}

}
