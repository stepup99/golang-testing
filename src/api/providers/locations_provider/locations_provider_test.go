package locations_provider

import (
	"net/http"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestGetCountryResetClientError(t *testing.T) {
	rest.StartMockupServer()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/AR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: 0,
	})
	// Execution``
	country, err := GetCountry("AR")
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid restclient error when getting country AR", err.Message)
	// Validation
}

func TestGetCountryNotFound(t *testing.T) {
	rest.StartMockupServer()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/AR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message":"Country not found", "error":"not_found", "status":404, "cause":[]}`,
	})
	country, err := GetCountry("AR")
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "Country not found", err.Message)
}

func TestGetCountryInvalidErrorInterface(t *testing.T) {
	rest.StartMockupServer()
	country, err := GetCountry("AR")
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid error interface when getting country AR", err.Message)
}

func TestGetCountryInvalidJsonResponse(t *testing.T) {
	rest.StartMockupServer()
	country, err := GetCountry("AR")
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "errror when trying to unmrshal country data for AR", err.Message)
}

func TestGetCountryNoError(t *testing.T) {
	rest.StartMockupServer()
	country, err := GetCountry("AR")
	assert.Nil(t, err)
	assert.NotNil(t, country)
	assert.EqualValues(t, "AR", country.Id)
	assert.EqualValues(t, "Argentina", country.Name)
	assert.EqualValues(t, "GMT-03:00", country.TimeZone)
	assert.EqualValues(t, 24, len(country.States))
}

// need to check how manny scenarios are there
// timeout
// error from api
// invalid error interface
// valid response
//          invalid json response
//          valid json response no error
