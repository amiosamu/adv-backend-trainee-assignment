package v1

import (
	"github.com/amiosamu/adv-backend-trainee-assignment/internal/entity"
	"github.com/amiosamu/adv-backend-trainee-assignment/internal/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
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
// @Tags advertisements
// @Accept json
// @Produce json
// @Success 201 {object} v1.advertisementsRoutes.create.response
// @Router /api/v1/advertisements/create [post]
func (r *advertisementRoutes) create(context *gin.Context) {
	id, err := r.advertisementService.CreateAdvertisement(context.Request.Context(), entity.Advertisement{})
	if err != nil {
		log.Fatal()
	}
	context.JSON(http.StatusCreated, map[string]interface{}{
		"Id": id,
	})
}
