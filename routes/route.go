package routes

import (
	"net/http"

	"github.com/devnura/pre-tets-devnura/config"
	_ "github.com/devnura/pre-tets-devnura/docs"
	"github.com/devnura/pre-tets-devnura/handler"
	"github.com/devnura/pre-tets-devnura/helper"
	_middleware "github.com/devnura/pre-tets-devnura/middleware"
	"github.com/devnura/pre-tets-devnura/repository"
	"github.com/devnura/pre-tets-devnura/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	echoSwagger "github.com/swaggo/echo-swagger"
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

		userRepo     repository.UserRepository     = repository.NewUserRepository(gormDB)
		questionRepo repository.QuestionRepository = repository.NewQuestionRepository(gormDB)
	)

	// service
	var (
		jwtService = service.NewJWTService()

		authService     = service.NewAuthService(userRepo)
		userService     = service.NewUserService(userRepo)
		questionService = service.NewQuestionService(questionRepo)
	)

	// handler
	var (
		authHandler     = handler.NewAuthHandler(authService, jwtService)
		userHandler     = handler.NewUserHandler(userService)
		questionHandler = handler.NewQuestionHandler(questionService)
	)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, helper.Response{
			Code:    200,
			Error:   []helper.EmptyObj{},
			Message: "Hello from the other side",
			Data:    []helper.EmptyObj{},
		})
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	v1 := e.Group("/api/v1")
	{
		v1.POST("/login", authHandler.Login)
		v1.GET("/profile", userHandler.Profile, _middleware.IsLoggedIn)
		v1.GET("/question", questionHandler.All, _middleware.IsLoggedIn)
		v1.GET("/question/:id", questionHandler.FindById, _middleware.IsLoggedIn)
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

}
