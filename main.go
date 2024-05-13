package main

import (
	"flag"
	config "golang_api/configs"
	"golang_api/routers"
	"net/http"
	"os"
	"strings"

	"golang_api/docs"
	"golang_api/migrations"

	"github.com/gin-gonic/gin"
)

var (
	ApiVersion  string
	Environment string
)

// @contact.name				API Support
// @contact.url				https://www.linkedin.com/in/yayang-suryana-308a5213a/
// @contact.email				yankzsoe@gmail.com
// @license.name				Source Code
// @license.url				https://github.com/yankzsoe/golang_api
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

	// Implement Custom Migrations
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
	docs.SwaggerInfo.Schemes = []string{"https", "http"}

	r.Run(":5001")
}

func GetVersion(ctx *gin.Context) {
	schema := "http"
	if strings.Contains(ctx.Request.Proto, "HTTPS") {
		schema = "https"
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":     "OK",
		"apiVersion": &ApiVersion,
		"message":    "Please visit " + schema + "://" + ctx.Request.Host + "/swagger/index.html for more documentation.",
	})
}
