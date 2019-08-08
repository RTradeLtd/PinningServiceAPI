package api

import (
	"github.com/RTradeLtd/PinningServiceAPI/api/models"
	"github.com/gin-gonic/gin"
)

func NewError(code int32, message string) *models.Error {
	return &models.Error{Code: code, Message: message}
}

func (a *API) handleError(c *gin.Context, code int32, errMsg string) {
	c.JSON(int(code), NewError(code, errMsg))
}
