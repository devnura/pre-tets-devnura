package routes

import (
	"net/http"

	"github.com/devnura/pre-tets-devnura/config"
	"github.com/devnura/pre-tets-devnura/handler"
	"github.com/devnura/pre-tets-devnura/repository"
	"github.com/devnura/pre-tets-devnura/service"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func SetupLogger(e *echo.Echo) {
	e.Logger.SetLevel(log.DEBUG)
}

func SetupMiddleware(e *echo.Echo) {
	// global or root middleware
	e.Use(middleware.RequestID())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

}

func SetupRoute(e *echo.Echo) {

	// repository
	var (
		gormDB = config.NewMysqlDB()

		userRepo repository.UserRepository = repository.NewUserRepository(gormDB)
	)

	// service
	var (
		authService = service.NewAuthService(userRepo)
		jwtService  = service.NewJWTService()
	)

	// handler
	var (
		authHandler = handler.NewAuthHandler(authService, jwtService)
	)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from another world!")
	})

	g := e.Group("/api/v1/auth")

	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &jwt.StandardClaims{},
		SigningKey: []byte("secret"),
	}))

	g.POST("/login", authHandler.Login)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

}
