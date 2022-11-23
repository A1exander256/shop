package handler

import (
	"net/http"
	"strconv"

	"github.com/alexander256/shop/models"
	"github.com/gin-gonic/gin"
)

type outputCreateUser struct {
	Id int `json:"id"`
}

func (h *Handler) createUser(c *gin.Context) {
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		h.responseError(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.service.User.Create(&input)
	if err != nil {
		h.log.Error(err)
		h.responseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	var result outputCreateUser
	result.Id = id
	c.JSON(http.StatusOK, result)
}

type outputUpdateUser struct {
	Status string `json:"status"`
}

func (h *Handler) updateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.responseError(c, http.StatusBadRequest, "invalid id param")
		return
	}
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		h.responseError(c, http.StatusBadRequest, err.Error())
		return
	}

	input.Uuid = id
	if err = h.service.User.Update(&input); err != nil {
		h.responseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	var result outputUpdateUser
	result.Status = "ok"
	c.JSON(http.StatusOK, result)
}

type outputDeleteUser struct {
	Status string `json:"status"`
}

func (h *Handler) deleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.responseError(c, http.StatusBadRequest, "invalid id param")
		return
	}
	if err := h.service.User.Delete(id); err != nil {
		h.responseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	var result outputDeleteUser
	result.Status = "ok"
	c.JSON(http.StatusOK, result)

}
