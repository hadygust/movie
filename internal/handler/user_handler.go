package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hadygust/movie/internal/dto"
	"github.com/hadygust/movie/internal/service"
)

type UserHandler struct {
	svc service.UserService
}

func NewUserHandler(svc service.UserService) UserHandler {
	return UserHandler{
		svc: svc,
	}
}

func (h *UserHandler) AllUser(c *gin.Context) {
	users, err := h.svc.AllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) Register(c *gin.Context) {
	var user dto.RegisterRequest

	err := c.ShouldBindBodyWithJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := h.svc.RegisterUser(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *UserHandler) Login(c *gin.Context) {
	var input dto.LoginRequest

	err := c.ShouldBindBodyWithJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := h.svc.Login(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
