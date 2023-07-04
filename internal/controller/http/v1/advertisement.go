package v1

import (
	"errors"
	"github.com/amiosamu/adv-backend-trainee-assignment/internal/entity"
	"github.com/amiosamu/adv-backend-trainee-assignment/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
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
type getAdvertisementResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Pictures    []string  `json:"pictures"`
	Price       int       `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
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
// @Description Get advertisement by Id
// @Tags advertisements
// @Accept json
// @Produce json
// @Success 201 {object} v1.advertisementsRoutes.get.response
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/advertisements/:id [get]

func (r *advertisementRoutes) get(context *gin.Context) {
	id := context.Param("id")

	adID, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid advertisement ID"})
		return
	}

	advertisement, err := r.advertisementService.GetAdvertisementById(context.Request.Context(), adID)
	if err != nil {
		if errors.Is(err, service.ErrAdvertisementNotFound) {
			context.JSON(http.StatusNotFound, gin.H{"error": "Advertisement not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	response := getAdvertisementResponse{
		ID:          advertisement.Id,
		Name:        advertisement.Name,
		Description: advertisement.Description,
		Pictures:    advertisement.Pictures,
		Price:       advertisement.Price,
		CreatedAt:   advertisement.CreatedAt,
	}

	context.JSON(http.StatusOK, response)
}
