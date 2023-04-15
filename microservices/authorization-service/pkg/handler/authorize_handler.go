package handler
import (
    "github.com/gin-gonic/gin"
)

func AuthorizeHandler(c *gin.Context) {
    c.String(200, "v1: %s %s", c.Request.Method, c.Request.URL.Path)
} 