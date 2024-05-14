
# GoEpexSpot


![alt text](static/logo400.png)

GoEpexSpot aims to be a comprehensive API client for fetching Day-ahead EPEX SPOT prices using Golang. 

Currently, EPEX SPOT and Nord Pool does not offer a free (or easily accesible) API for non-commercial uses.
However, some stakeholders do provide access. This repository aims to highlight such resources, starting with the French Transmission System Operator (RTE) API for real-time data access.

This repository facilitates price-driven demand response, thereby encouraging environmentally friendly behavior as prices frequently correlate with CO2 marginal emissions.



| Features                                              | Status      |
|-------------------------------------------------------|-------------|
| Regroup different country Go API Client               | ✅          |
| Fetch real-time or D-1 electricity price data         | ✅          |
| Provide steps to access data from different providers | ✅          |
| Enable history fetching                               | ❌          |
| Store historical data                                 | ❌          |


| Country              | Status      |
|----------------------|-------------|
| Belgium              | 🔜          |
| France               | ✅          |
| Germany              | 🔜          |
| Netherland           | 🔜          |




## Description

GoEpexSpot is designed to unlock the potential of Day-ahead EPEX spot prices for electricity through a unique Go-based API client. In the absence of a freely accessible electricty price API from EPEX SPOT or NordPool for non-commercial purposes, this repository shines a spotlight on alternative sources, beginning with access via the French Transmission System Operator (RTE) API for timely data retrieval. By enabling a price-driven demand response, GoEpexSpot not only facilitates financially smarter energy consumption decisions but also promotes eco-friendly practices as energy price tend to be linked with CO2 emissions.


## Prerequisite

### French Provider

To use this repository, please ensure you meet the following requirements:

- **Account on RTE API Website**: You must have an account to access the RTE API. [Sign up or log in here](https://data.rte-france.com).
- **Subscription to the "Wholesale Market" API**: Access requires a subscription to this specific API. Without it, you will encounter a 403 error. [Subscribe here](https://data.rte-france.com/catalog/-/api/market/Wholesale-Market/v2.0#).

## Usage Example

```go
basicAuthUsername := "YOUR_RTE_API_BASIC_AUTH_USERNAME"
basicAuthPassword := "YOUR_RTE_API_BASIC_AUTH_PASSWORD"

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
```

## Contribution

Currently, this repository provides access exclusively to data for France. It is, however, designed to support additional countries in the future. Contributions to expand its scope and capabilities are highly encouraged.

## Usage Notes

### RTE API

The data accessed through the Wholesale Market API is exclusively owned by EPEX SPOT SE, its subsidiaries, and Nord Pool AS. While users are authorized to download, print, and use the data for internal and personal purposes (with proper source and ownership citation), the following uses are strictly prohibited without written consent from EPEX SPOT SE and Nord Pool AS:

- Any form of commercial use.
- The creation of any financial instrument or reference index for external use or the benefit of third parties.
- Any copying, distribution, marketing, exploitation, or use benefiting third parties.