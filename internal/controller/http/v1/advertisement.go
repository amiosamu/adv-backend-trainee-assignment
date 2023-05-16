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
	c.POST("/create", r.create)
}

// @Summary Create advertisement
// @Description Create advertisement
func (r *advertisementRoutes) create(context *gin.Context) {
	id, err := r.advertisementService.CreateAdvertisement(context.Request().Context())
}
