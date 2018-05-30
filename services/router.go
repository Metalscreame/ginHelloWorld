package services

import (
	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/services/singleArticle"
	"github.com/heroku/go-getting-started/services/showIndex"
	"os"
	"log"
)

var router *gin.Engine

func InitializeRouter() {

	//Used for heroku
	port := os.Getenv("PORT")

	//Uncomment for local machine
	//port="8080"
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router = gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/static", "static")


	router.GET("/", showIndex.ShowIndexPage)
	router.GET("/article/view/:article_id", singleArticle.GetArticle)

	router.Run(":" + port)
}