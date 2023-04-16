package handler

import (
	"go/pkg/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) TestHandler(c *gin.Context) {
	if user, err := jwt.ExtractTokenUser(c); err != nil {
		c.JSON(http.StatusBadRequest, "Error")
	} else {
		c.JSON(http.StatusAccepted, user)
	}
}
