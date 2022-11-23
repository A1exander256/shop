package handler

import "github.com/gin-gonic/gin"

type messageError struct {
	Message string `json:"message"`
}

func (h *Handler) responseError(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, messageError{Message: message})
}
