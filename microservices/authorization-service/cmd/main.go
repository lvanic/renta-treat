package main

import (
	"go/pkg/config"
	"go/pkg/model"
	"go/pkg/router"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	viper.SetConfigFile("../config/.env")
	viper.ReadInConfig()

	config := config.NewConfig()
	db, err := gorm.Open(mysql.Open(config.DatabaseURL), &gorm.Config{})

	if err != nil {
		panic("Database open error")
	}

	db.AutoMigrate(&model.User{})
	r := gin.Default()
	router.RegisterRoutes(r, db)

	log.Fatal(r.Run(config.Port))

}
