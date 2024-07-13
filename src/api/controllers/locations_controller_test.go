package controllers

import (
	"encoding/json"
	"flag"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stepup99/golang-testing/src/api/domain/locations"
	"github.com/stepup99/golang-testing/src/api/utils/errors"
	"github.com/stretchr/testify/assert"
)

var getCountryFunc func(countryId string) (*locations.Country, *errors.ApiError)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestGetCountryNotFound(t *testing.T) {
	flag.Parse()
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/AR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message":"Country not found", "error":"not_found", "status":404, "cause":[]}`,
	})
	// take reponse recorders
	response := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(response)
	// create a request
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params{
		{Key: "country_id", Value: "AR"},
	}

	GetCountry(c)

	assert.EqualValues(t, http.StatusNotFound, response.Code)
	var apiErr errors.ApiError
	err := json.Unmarshal(response.Body.Bytes(), &apiErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusNotFound, apiErr.Status)
	assert.EqualValues(t, "Country not found", apiErr.Message)
}
