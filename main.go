package main

import (
	"flag"
	config "golang_api/configs"
	"golang_api/routers"
	"log"
	"net/http"
	"os"

	"golang_api/app/migrations"
	"golang_api/docs"

	"github.com/gin-gonic/gin"
)

var (
	ApiVersion  string
	Environment string
)

// @contact.name				API Support
// @contact.url				http://www.swagger.io/support
// @contact.email				support@swagger.io
// @license.name				Apache 2.0
// @license.url				http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	// Setup GIN in release mode
	gin.SetMode(gin.ReleaseMode)

	// Read flag from -ldflag
	flag.Parse()
	// Get Instance for JWT Configuration
	// this process read file only once
	// or called Singleton pattern
	config.GetJWTConfigurationInstance()

	// Initialize connection to Database
	config.InitDB()

	// Implement Custom Migration
	if err := migrations.Apply(); err != nil {
		panic(err)
	}

	// Initialize router
	r := routers.SetupRouter()

	r.GET("/", GetVersion)

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger User API"
	docs.SwaggerInfo.Description = "This is a sample Swagger in Golang with GIN Framework."
	docs.SwaggerInfo.Version = "1.0"
	if Environment == "PRODUCTION" {
		docs.SwaggerInfo.Host = os.Getenv("hostSwagger")
	} else {
		docs.SwaggerInfo.Host = "localhost:5001/api"
	}
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"https"}

	if Environment == "PRODUCTION" {
		r.Run(":5001")
	} else {
		log.Fatal(http.ListenAndServeTLS(":5001", "server.crt", "server.key", r))
	}
}

func GetVersion(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":     "OK",
		"apiVersion": &ApiVersion,
		"message":    "Please visit https://golang-api-6ej0.onrender.com/swagger/index.html for more documentation.",
	})
}
