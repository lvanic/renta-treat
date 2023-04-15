package main

import (
	"log"
	// "net/http"
	// "gorm.io/driver/mysql"
	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
)

func main() {
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local";
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{});

	// if err != nil {
	// 	db.First();
	// }
	router := gin.Default()

	router.GET("/authorize", AuthorizeHandler)

	log.Fatal(router.Run(":8080"))
}