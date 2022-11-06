package main

import (
	"github.com/devnura/pre-tets-devnura/routes"
	"github.com/labstack/echo/v4"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct {
}

func main() {
	var server = echo.New()

	routes.SetupLogger(server)
	// routes.SetupGlobalErrorHandler(server)
	routes.SetupMiddleware(server)
	routes.SetupRoute(server)

	server.Logger.Fatal(server.Start(":8080"))

}
