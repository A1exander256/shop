package handler

import (
	_ "github.com/alexander256/shop/docs"
	"github.com/alexander256/shop/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	users := router.Group("/users")
	{
		users.POST("/", h.createUser)
		users.PUT("/:id", h.updateUser)
		users.DELETE("/:id", h.deleteUser)
		users.GET("/:id", h.getUserById)
		users.GET("/", h.getAllUsers)
	}

	orders := router.Group("/orders")
	{
		orders.POST("/", h.createOrder)
	}

	products := router.Group("/products")
	{
		products.POST("/", h.createProduct)
	}
	return router
}
