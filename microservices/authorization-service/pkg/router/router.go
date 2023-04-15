package router

import (
	"go/pkg/handler"
	"go/pkg/repository"
	"go/pkg/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(*userRepository)

	h := handler.Handler{
		Db:          db,
		UserService: userService,
	}
	router.GET("/test", h.TestHandler)
	router.POST("/authorization", h.AuthorizeHandler)
	router.POST("/registration", h.RegistrationHandler)
}
