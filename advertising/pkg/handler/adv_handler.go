package handler

import (
	"context"
	"github.com/amiosamu/adv-backend-trainee-assignment/advertising/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

var ctx context.Context

func (h *Handler) createAdv(c *gin.Context) {
	var input model.Advertising
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Advertising.Create(ctx, input)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

func (h *Handler) getAllAdv(c *gin.Context) {
	advId := uuid.Must(uuid.FromBytes([]byte(c.Param("id"))))
	allAdv, err := h.services.Advertising.GetAll(ctx, advId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, allAdv)
}

func (h *Handler) getById(c *gin.Context) {
	advId := uuid.Must(uuid.FromBytes([]byte(c.Param("id"))))
	adv, err := h.services.Advertising.GetByID(ctx, advId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, adv)
}
