package v1

import (
	"github.com/amiosamu/adv-backend-trainee-assignment/internal/entity"
	"github.com/amiosamu/adv-backend-trainee-assignment/internal/service"

	"github.com/gin-gonic/gin"
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
	var input entity.Advertisement
	id, err := r.advertisementService.CreateAdvertisement(context.Request.Context(), input)
	if err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}
	context.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
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
	var input getByIdInput

	if err := context.BindJSON(input); err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}
	advertisement, err := r.advertisementService.GetAdvertisementById(context.Request.Context(), input.Id)
	if err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}
	context.JSON(http.StatusOK, advertisement)
}
