package handler

import (
	"go/pkg/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	Db          *gorm.DB
	UserService *service.UserService
}

func (h *Handler) AuthorizeHandler(c *gin.Context) {
	c.String(200, "v1: %s %s", c.Request.Method, c.Request.URL.Path)
}
