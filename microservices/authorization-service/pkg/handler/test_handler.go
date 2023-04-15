package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) TestHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(400, "id must be a number")
	}
	user, err := h.UserService.GetUserById(id)
	if err != nil {
		c.JSON(400, "User not found")
	}
	c.JSON(200, &user)
}
