package services

import (
	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/services/singleArticle"
	"github.com/heroku/go-getting-started/services/showIndex"
)

// Handle routes
func InitializeRoutes(router *gin.Engine) {
	router.GET("/", showIndex.ShowIndexPage)
	router.GET("/article/view/:article_id", singleArticle.GetArticle)
}
