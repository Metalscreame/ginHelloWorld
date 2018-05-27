package services

import (
	"github.com/heroku/go-getting-started/services/showIndex"
	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/services/singleArticle"
)

func InitializeRoutes(router *gin.Engine) {
	// Handle the index route
	router.GET("/", showIndex.ShowIndexPage)
	router.GET("/article/view/:article_id", singleArticle.GetArticle)
}
