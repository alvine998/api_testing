package main

import (
	"api_testing/handler"
	"api_testing/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/api_testing?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Error")
	}
	fmt.Println("Database Connect Successfully")
	db.AutoMigrate(&user.User{})

	userRepository := user.NewRepository(db)
	userService := user.NewServiceUser(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", userHandler.HelloWorld)

	router.Run()

}
