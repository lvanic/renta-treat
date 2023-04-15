package main

import (
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

	port := viper.Get("PORT").(string)
	dsn := viper.Get("DatabaseURL").(string)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Database open error")
	}

	db.AutoMigrate(&model.User{})
	r := gin.Default()
	router.RegisterRoutes(r, db)

	log.Fatal(r.Run(port))

}
