package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/timebis/goEpexSpot"
)

func main() {

	err := godotenv.Load("/home/thomas/git/timebis/GoEpexSpot/.env")
	if err != nil {
		fmt.Printf("Error loading .env file: %+v", err)
		panic(err)
	}

	basicAuthUsername := os.Getenv("RTE_API_BASIC_AUTH_USERNAME")
	basicAuthPassword := os.Getenv("RTE_API_BASIC_AUTH_PASSWORD")

	auth := goEpexSpot.AuthOptions{
		Username: basicAuthUsername,
		Password: basicAuthPassword,
	}

	fmt.Printf("Basic Auth Username: %s\n", basicAuthUsername)

	// Optionally, you can use your independently-fetched bearer token
	// bearerToken := os.Getenv("BEARER_TOKEN")
	// auth := AuthOptions{
	// 	BearerToken: bearerToken,
	// }

	epexSpotDayAhead, err := goEpexSpot.GetEpexSpot(goEpexSpot.France, auth)
	if err != nil {
		log.Fatalf("Error while fetching data: %+v\n", err)
	} else {
		fmt.Printf("Market Data: %+v\n", epexSpotDayAhead)
	}
}
