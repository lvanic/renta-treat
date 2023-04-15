package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) RegistrationHandler(c *gin.Context) {
	c.String(200, "Register: %s %s", c.Request.Method, c.Request.URL.Path)
}
