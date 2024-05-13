package main

import (
	"fmt"
	"goEpexSpot"
	"log"
	"os"
)

func main() {

	basicAuthUsername := os.Getenv("BASIC_AUTH_USERNAME")
	basicAuthPassword := os.Getenv("BASIC_AUTH_PASSWORD")
	fmt.Printf("token : %v \n", basicAuthUsername)

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
