package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	userCtx = "uuid"
)

func getAdvId(c *gin.Context) (uuid.UUID, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return uuid.Nil, errors.New("error getting user")
	}
	idInt, ok := id.(uuid.UUID)
	if !ok {
		return uuid.Nil, errors.New("user id is of invalid type")
	}
	return idInt, nil
}
