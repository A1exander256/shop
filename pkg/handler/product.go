package handler

import (
	"net/http"
	"strconv"

	"github.com/alexander256/shop/models"
	"github.com/gin-gonic/gin"
)

type outputCreateProduct struct {
	Id int `json:"id"`
}

func (h *Handler) createProduct(c *gin.Context) {
	var input models.Product
	if err := c.BindJSON(&input); err != nil {
		h.responseError(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.service.Product.Create(&input)
	if err != nil {
		h.log.Error(err)
		h.responseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	var result outputCreateProduct
	result.Id = id
	c.JSON(http.StatusOK, result)
}

type outputUpdateProduct struct {
	Status string `json:"status"`
}

func (h *Handler) updateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.responseError(c, http.StatusBadRequest, "invalid id param")
		return
	}
	var input models.Product
	if err := c.BindJSON(&input); err != nil {
		h.responseError(c, http.StatusBadRequest, err.Error())
		return
	}
	input.Id = id
	if err = h.service.Product.Update(&input); err != nil {
		h.responseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	var result outputUpdateProduct
	result.Status = "ok"
	c.JSON(http.StatusOK, result)
}

type outputDeleteProduct struct {
	Status string `json:"status"`
}

func (h *Handler) deleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.responseError(c, http.StatusBadRequest, "invalid id param")
		return
	}
	if err := h.service.Product.Delete(id); err != nil {
		h.responseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	var result outputDeleteProduct
	result.Status = "ok"
	c.JSON(http.StatusOK, result)
}
