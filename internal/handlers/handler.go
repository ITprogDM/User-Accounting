package handlers

import (
	"UchetUsers/internal/middleware"
	"UchetUsers/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type UserHandler struct {
	service services.UserService
	logger  *logrus.Logger
}

func NewUserHandler(service services.UserService, logger *logrus.Logger) *UserHandler {
	return &UserHandler{
		service: service,
		logger:  logger,
	}
}

func (h *UserHandler) InitRoutes(logger *logrus.Logger) *gin.Engine {
	rout := gin.New()

	rout.Use(middleware.RequestLoggerMiddleware(logger))

	rout.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	user := rout.Group("/users")
	{
		user.POST("", h.CreateUserHandler)
		user.GET("/:id", h.GetUserHandler)
		user.PUT("/:id", h.UpdateUserHandler)
		user.DELETE("/:id", h.DeleteUserHandler)
	}

	return rout
}
