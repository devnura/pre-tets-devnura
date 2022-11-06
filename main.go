package main

import (
	"github.com/devnura/pre-tets-devnura/routes"
	"github.com/labstack/echo/v4"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization

func main() {
	var server = echo.New()

	routes.SetupLogger(server)
	// routes.SetupGlobalErrorHandler(server)
	routes.SetupMiddleware(server)
	routes.SetupRoute(server)

	server.Logger.Fatal(server.Start(":8080"))

}
