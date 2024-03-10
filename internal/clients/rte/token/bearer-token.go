package token

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

const (
	bearerApiBaseUrl = "https://digital.iservices.rte-france.com/token/oauth/"
)

// var basicAuthUsername string
// var basicAuthPassword string

type tokenResp struct {
	AccessToken string `json:"access_token"`
	Type        string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func FetchBearerToken(basicAuthUsername string, basicAuthPassword string) (string, error) {

	req, err := http.NewRequest("GET", bearerApiBaseUrl, nil)
	if err != nil {
		return "", errors.Wrap(err, "creating request :")
	}

	auth := base64.StdEncoding.EncodeToString([]byte(basicAuthUsername + ":" + basicAuthPassword))
	req.Header.Add("Authorization", "Basic "+auth)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	if res.StatusCode != http.StatusOK {
		return "", errors.Errorf("status %v when fetching token", res.StatusCode)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", errors.Wrap(err, "reading token response")
	}

	var apiResp tokenResp
	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		return "", errors.Wrap(err, "unmarshalling token response")
	}

	return apiResp.AccessToken, nil
}
