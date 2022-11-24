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

// @Summary Create user
// @Tags users
// @Description create user
// @ID create-user
// @Accept  json
// @Produce  json
// @Param input body models.User true "user info"
// @Success 200 {object} outputCreateUser
// @Failure 400 {object} messageError
// @Router /users [post]
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

// @Summary Update user
// @Tags users
// @Description update user
// @ID update-user
// @Accept  json
// @Produce  json
// @Param input body models.User true "user info"
// @Success 200 {object} outputUpdateUser
// @Failure 400 {object} messageError
// @Router /users/:id [put]
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

// @Summary Delete user
// @Tags users
// @Description delete user
// @ID create-user
// @Accept  json
// @Produce  json
// @Success 200 {object} outputDeleteUser
// @Failure 400 {object} messageError
// @Router /users/:id [delete]
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

// @Summary Get User By Id
// @Tags users
// @Description get user by id
// @ID get-user-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User
// @Failure 400 {object} messageError
// @Router /users/:id [get]
func (h *Handler) getUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.responseError(c, http.StatusBadRequest, "invalid id param")
		return
	}

	user, err := h.service.User.GetById(id)
	if err != nil {
		h.responseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary Get all Users
// @Tags users
// @Description get all users
// @ID get-all-users
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.User
// @Failure 400 {object} messageError
// @Router /users [get]
func (h *Handler) getAllUsers(c *gin.Context) {
	users, err := h.service.User.GetAll()
	if err != nil {
		h.responseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, users)
}
