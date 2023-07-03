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
	c.GET("/:id", r.get)

}

// @Summary Create advertisement
// @Description Create advertisement
// @Tags advertisements
// @Accept json
// @Produce json
// @Success 201 {object} v1.advertisementsRoutes.create.response
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/advertisements/create [post]

func (r *advertisementRoutes) create(context *gin.Context) {

}

type getByIdInput struct {
	Id int `json:"id" validate:"required"`
}

// @Summary Get advertisement
// @Description Get  advertisement by Id
// @Tags advertisements
// @Accept json
// @Produce json
// @Success 201 {object} v1.advertisementsRoutes.get.response
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/advertisements/:id [post]
func (r *advertisementRoutes) get(context *gin.Context) {

}
