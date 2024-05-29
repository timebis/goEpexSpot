package main

import (
	"fmt"
	"log"

	"github.com/timebis/goEpexSpot"
)

func main() {
	basicAuthUsername := "YOUR_BASIC_AUTH_USERNAME"
	basicAuthPassword := "YOUR_BASIC_AUTH_PASSWORD"

	auth := goEpexSpot.AuthOptions{
		Username: basicAuthUsername,
		Password: basicAuthPassword,
	}

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
