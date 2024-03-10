package goEpexSpot

import (
	"context"
	"goEpexSpot/internal/clients/rte/swagger"
	"goEpexSpot/internal/clients/rte/token"
	"net/http"
	"time"

	errors "github.com/pkg/errors"
)

type EpexSpotDayAhead struct {
	// Start date for data requested
	StartDate time.Time `json:"start_date,omitempty"`
	// End date for requested data
	EndDate time.Time `json:"end_date,omitempty"`
	// Date updated
	UpdatedDate time.Time        `json:"updated_date,omitempty"`
	Values      []EpexSpotResult `json:"values,omitempty"`
}
type EpexSpotResult struct {
	// Start time interval
	StartDate time.Time `json:"start_date,omitempty"`
	// End time interval
	EndDate time.Time `json:"end_date,omitempty"`
	// Volume of the electricity market (in MW)
	Value_mw float32 `json:"value,omitempty"`
	// Price (in €/MWh)
	Price_eur_per_mwh float32 `json:"price,omitempty"`
}

// AuthOptions holds authentication details
type AuthOptions struct {
	Username    string
	Password    string
	BearerToken string
}

type Country string

const (
	France Country = "FR"
)

func GetEpexSpot(countryCode Country, authOpts ...AuthOptions) (data EpexSpotDayAhead, err error) {
	var auth AuthOptions

	if len(authOpts) > 0 {
		auth = authOpts[0]
	}

	switch countryCode {
	case France:
		return rte(auth)
	default:
		return data, errors.New("country code not supported")
	}
}

func rte(auth AuthOptions) (epexSpotDayAhead EpexSpotDayAhead, err error) {

	// check if the auth options are valid given rte API
	if auth.BearerToken == "" && (auth.Username == "" || auth.Password == "") {
		return EpexSpotDayAhead{}, errors.New("BearerToken or {Username, Password} must be provided for french API")

	}

	var bearerToken = auth.BearerToken

	// Fetch the bearer token if not provided
	if auth.BearerToken == "" {
		basicAuthUsername := auth.Username
		basicAuthPassword := auth.Password
		bearerToken, err = token.FetchBearerToken(basicAuthUsername, basicAuthPassword)
		if err != nil {
			return epexSpotDayAhead, errors.Wrap(err, "Error while fetching bearer token")
		}

	}

	// Create a new API client configuration
	cfg := swagger.NewConfiguration()
	cfg.HTTPClient = &http.Client{}
	cfg.AddDefaultHeader("Authorization", "Bearer "+bearerToken) // Add the Authorization header with the Bearer token

	client := swagger.NewAPIClient(cfg)

	frenchData, resp, err := client.DefaultApi.GetFrancePowerExchanges(context.Background())
	_ = resp
	if err != nil {
		return epexSpotDayAhead, errors.Wrap(err, "Error while fetching data from RTE API")
	}
	// cast the response to the expected format
	epexSpotDayAhead = convertFenchResultToStandartResult(frenchData)

	return epexSpotDayAhead, err
}

func convertFenchResultToStandartResult(frenchData swagger.RespData) (standartData EpexSpotDayAhead) {
	standartData.StartDate = frenchData.FrancePowerExchange[0].StartDate
	standartData.EndDate = frenchData.FrancePowerExchange[0].EndDate
	standartData.UpdatedDate = frenchData.FrancePowerExchange[0].UpdatedDate

	for _, value := range frenchData.FrancePowerExchange[0].Values {
		standartData.Values = append(standartData.Values, EpexSpotResult{
			StartDate:         value.StartDate,
			EndDate:           value.EndDate,
			Value_mw:          value.Value,
			Price_eur_per_mwh: value.Price,
		})
	}

	return standartData
}
