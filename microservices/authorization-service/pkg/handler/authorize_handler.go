package handler

import (
	"go/pkg/jwt"
	"go/pkg/model"
	"go/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	Db          *gorm.DB
	UserService *service.UserService
}

type userAuthorize struct {
	Email    string
	Password string
}

func (h *Handler) AuthorizeHandler(c *gin.Context) {
	var userAuth userAuthorize

	if err := c.ShouldBindJSON(&userAuth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Email:    userAuth.Email,
		Password: userAuth.Password,
	}

	if userHandler, err := h.UserService.AuthorizeUser(&user); err != nil {
		c.JSON(http.StatusNotFound, err.Error())
	} else {
		jwtToken, _ := jwt.GenerateToken(userHandler)
		userResponse := struct {
			Name        string
			Email       string
			AcsessToken string
		}{
			Name:        userHandler.Name,
			Email:       userHandler.Email,
			AcsessToken: jwtToken,
		}
		c.JSON(http.StatusAccepted, userResponse)
	}
}
