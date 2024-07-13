package locations_provider

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stepup99/golang-testing/src/api/domain/locations"
	"github.com/stepup99/golang-testing/src/api/utils/errors"
)

const (
	urlGetCountry = "https://api.mercadolibre.com/countries/%s"
)

func GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	fmt.Println("countryId >>>>>>>.......-----------   ", countryId)
	response := rest.Get(fmt.Sprintf(urlGetCountry, countryId))
	fmt.Println("response >>>>>>>>>..........=--------------- ", response)
	fmt.Println("response.Response >>>>>>>>---------------", response.Response)
	if response == nil || response.Response == nil {
		fmt.Println("inside first if --------------->>>>>>>>>>")
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("invalid restclient error when getting country %s", countryId),
		}
	}

	if response.StatusCode > 299 {
		fmt.Println("inside second if --------------->>>>>>>>>>")
		var apiErr errors.ApiError
		if err := json.Unmarshal(response.Bytes(), &apiErr); err != nil {
			return nil, &errors.ApiError{
				Status:  http.StatusInternalServerError,
				Message: fmt.Sprintf("invalid error response when getting country %s", countryId),
			}
		}
		return nil, &apiErr
	}

	var result locations.Country
	if err := json.Unmarshal(response.Bytes(), &result); err != nil {
		fmt.Println("inside third if --------------->>>>>>>>>>")
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("error when trying to unmarshal country for %s", countryId),
		}
	}
	return &result, nil
}

// timeout
// error from api
// invalid error interface
// valid response
//          invalid json response
//          valid json response no error
