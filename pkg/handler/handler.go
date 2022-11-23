package handler

import (
	"github.com/alexander256/shop/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	log     *logrus.Logger
	service *service.Service
}

func NewHandler(service *service.Service, log *logrus.Logger) *Handler {
	return &Handler{
		log:     log,
		service: service,
	}
}

func (h *Handler) InitRoutes(ginMode string) *gin.Engine {
	gin.SetMode(ginMode)
	router := gin.New()

	user := router.Group("/user")
	{
		user.POST("/", h.createUser)
		user.PUT("/:id", h.updateUser)
		user.DELETE("/:id", h.deleteUser)
	}
	return router
}
