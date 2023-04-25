package handler

import (
	"github.com/amiosamu/adv-backend-trainee-assignment/advertising/pkg/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	adv := router.Group("/adv")
	{
		adv.POST("/", h.createAdv)
		adv.GET("/:id", h.getById)
		adv.GET("/", h.getAllAdv)
	}
	return router
}
