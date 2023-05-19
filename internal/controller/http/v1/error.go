package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func newErrorResponse(c *gin.Context, errStatus int, message string) {
	err := errors.New(message)
	_, ok := err.(*gin.Error)
	if !ok {
		report := gin.H{
			"error": err.Error(),
		}
		c.JSON(errStatus, report)
	}
	c.Error(errors.New("internal server error"))
}
