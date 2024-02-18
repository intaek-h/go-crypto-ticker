package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/intaek-h/go-crypto-ticker/datatypes"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) {
	if len(currency) != 3 {
		return nil, fmt.Errorf("currency must be 3 characters. received: %v", len(currency))
	}
	upperCur := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiUrl, upperCur))
	if err != nil {
		return nil, err
	}
	var response CEXResponse
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("status code received: %v", res.StatusCode)
	}
	rate := &datatypes.Rate{Currency: upperCur, Price: response.Bid}
	return rate, nil
}
