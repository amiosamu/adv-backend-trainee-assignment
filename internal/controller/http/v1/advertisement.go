package v1

import (
	"errors"
	_ "github.com/amiosamu/adv-backend-trainee-assignment/docs"
	"github.com/amiosamu/adv-backend-trainee-assignment/internal/entity"
	"github.com/amiosamu/adv-backend-trainee-assignment/internal/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
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
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Pictures    []string `json:"pictures"`
	Price       int      `json:"price"`
	CreatedAt   string   `json:"created_at"`
}

type createAdvertisementResponse struct {
	ID   int `json:"id"`
	Code int `json:"code"`
}

func newAdvertisementRoutes(c *gin.RouterGroup, advertisementService service.Advertisement) {
	r := &advertisementRoutes{
		advertisementService: advertisementService,
	}
	c.POST("/create", r.create)
	c.GET("/:id", r.get)
	c.GET("/", r.getAll)
}

// @Summary Create advertisement
// @Description Create advertisement
// @Tags advertisements
// @Accept json
// @Produce json
// @Success 201 {object} v1.advertisementRoutes.create.createAdvertisementResponse
// @Failure 400 {object} createAdvertisementResponse
// @Failure 500 {object} createAdvertisementResponse
// @Failure default {object} createAdvertisementResponse
// @Router /api/v1/advertisements/create [post]
func (r *advertisementRoutes) create(ctx *gin.Context) {
	var request createAdvertisementRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	advertisement := &entity.Advertisement{
		Name:        request.Name,
		Description: request.Description,
		Pictures:    request.Pictures,
		Price:       request.Price,
	}

	id, err := r.advertisementService.CreateAdvertisement(ctx.Request.Context(), advertisement)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := createAdvertisementResponse{
		ID:   id,
		Code: http.StatusCreated,
	}
	ctx.JSON(http.StatusCreated, response)
}

// @Summary Get advertisement
// @Description Get advertisement by ID
// @Tags advertisements
// @Accept json
// @Produce json
// @Param id path int true "Advertisement ID"
// @Success 200 {object} getAdvertisementResponse
// @Failure 400 {object} getAdvertisementResponse
// @Failure 404 {object} getAdvertisementResponse
// @Failure 500 {object} getAdvertisementResponse
// @Router /api/v1/advertisements/{id} [get]
func (r *advertisementRoutes) get(ctx *gin.Context) {
	id := ctx.Param("id")

	adID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid advertisement ID"})
		return
	}

	advertisement, err := r.advertisementService.GetAdvertisementById(ctx.Request.Context(), adID)
	if err != nil {
		log.Printf("failed to get advertisement by id: %v", err.Error())
		if errors.Is(err, service.ErrAdvertisementNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Advertisement not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	response := getAdvertisementResponse{
		ID:          advertisement.Id,
		Name:        advertisement.Name,
		Description: advertisement.Description,
		Pictures:    advertisement.Pictures,
		Price:       advertisement.Price,
		CreatedAt:   advertisement.CreatedAt.Format("January 2, 2006 15:04:05"),
	}

	ctx.JSON(http.StatusOK, response)
}

// @Summary Get advertisement
// @Description Get all advertisements
// @Tags advertisements
// @Produce json
// @Success 200 {object} v1.advertisementRoutes.getAll.response
// @Router /api/v1/advertisements/ [get]
func (r *advertisementRoutes) getAll(ctx *gin.Context) {
	advertisements, err := r.advertisementService.GetAdvertisements(ctx.Request.Context())
	if err != nil {
		log.Printf("failed to get advertisements: %v", err)

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	response := make([]getAdvertisementResponse, len(advertisements))
	for i, ad := range advertisements {
		response[i] = getAdvertisementResponse{
			ID:          ad.Id,
			Name:        ad.Name,
			Description: ad.Description,
			Pictures:    ad.Pictures,
			Price:       ad.Price,
			CreatedAt:   ad.CreatedAt.Format("January 2, 2006 15:04:05"),
		}
	}

	ctx.JSON(http.StatusOK, response)
}
