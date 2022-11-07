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
		answerRepo   repository.AnswerRepository   = repository.NewAnswerRepository(gormDB)
	)

	// service
	var (
		jwtService = service.NewJWTService()

		authService     = service.NewAuthService(userRepo)
		userService     = service.NewUserService(userRepo)
		questionService = service.NewQuestionService(questionRepo)
		answerService   = service.NewAnswerService(answerRepo)
	)

	// handler
	var (
		authHandler     = handler.NewAuthHandler(authService, jwtService)
		userHandler     = handler.NewUserHandler(userService)
		questionHandler = handler.NewQuestionHandler(questionService)
		answerHandler   = handler.NewAnswerHandler(answerService)
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

		v1.POST("/question", questionHandler.Insert, _middleware.IsLoggedIn)
		v1.GET("/question", questionHandler.All, _middleware.IsLoggedIn)
		v1.GET("/question/:id", questionHandler.FindById, _middleware.IsLoggedIn)
		v1.GET("/question/:id", questionHandler.FindAnswer, _middleware.IsLoggedIn)
		v1.PUT("/question/answer/:id", questionHandler.Update, _middleware.IsLoggedIn)
		v1.DELETE("/question/:id", questionHandler.Delete, _middleware.IsLoggedIn)

		v1.POST("/answer", answerHandler.InsertAnswer, _middleware.IsLoggedIn)
		v1.GET("/answer", answerHandler.AllAnswer, _middleware.IsLoggedIn)
		v1.GET("/answer/:id", answerHandler.FindAnswerById, _middleware.IsLoggedIn)
		v1.PUT("/answer/:id", answerHandler.UpdateAnswer, _middleware.IsLoggedIn)
		v1.DELETE("/answer/:id", answerHandler.DeleteAnswer, _middleware.IsLoggedIn)
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

}
