package v1

import (
	"github.com/amiosamu/adv-backend-trainee-assignment/internal/service"
	"github.com/gin-gonic/gin"
)

type advertisementRoutes struct {
	advertisementService service.Advertisement
}

func newAdvertisementRoutes(c *gin.RouterGroup, advertisementService service.Advertisement) {
	r := &advertisementRoutes{
		advertisementService: advertisementService,
	}
	c.POST("/")
}
