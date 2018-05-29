package main

import (
	"log"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"os"
	"github.com/metalscreame/ginHelloWorld/services"
)

var router *gin.Engine


func main() {
	port := os.Getenv("PORT")
	//port="8080"
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router = gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/static", "static")

	services.InitializeRoutes(router)
	router.Run(":" + port)
}

