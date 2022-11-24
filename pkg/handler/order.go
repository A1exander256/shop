package handler

import (
	"net/http"

	"github.com/alexander256/shop/models"
	"github.com/gin-gonic/gin"
)

type outputCreateOrder struct {
	Id int `json:"id"`
}

func (h *Handler) createOrder(c *gin.Context) {
	var input models.Order
	if err := c.BindJSON(&input); err != nil {
		h.responseError(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.service.Order.Create(&input)
	if err != nil {
		h.log.Error(err)
		h.responseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	var result outputCreateOrder
	result.Id = id
	c.JSON(http.StatusOK, result)
}
