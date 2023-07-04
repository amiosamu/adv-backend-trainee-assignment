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

type createAdvertisementRequest struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Pictures    []string `json:"pictures" binding:"required"`
	Price       int      `json:"price" binding:"required"`
}

type createAdvertisementResponse struct {
	ID int `json:"id"`
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
	var request createAdvertisementRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	advertisement := &entity.Advertisement{
		Name:        request.Name,
		Description: request.Description,
		Pictures:    request.Pictures,
		Price:       request.Price,
	}

	id, err := r.advertisementService.CreateAdvertisement(context.Request.Context(), advertisement)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := createAdvertisementResponse{ID: id}
	context.JSON(http.StatusCreated, response)
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
