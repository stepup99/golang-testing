package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stepup99/golang-testing/src/api/app"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}
func main() {
	// country, err := locations_provider.GetCountry("AR")
	// fmt.Println(err)
	// fmt.Println(country)
	// this is simple server getting created
	app.StartApp()
}
