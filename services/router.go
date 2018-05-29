package services

import (
	"github.com/gin-gonic/gin"
	"github.com/metalscreame/ginHelloWorld/services/singleArticle"
	"github.com/metalscreame/ginHelloWorld/services/showIndex"
)

// Handle routes
func InitializeRoutes(router *gin.Engine) {
	router.GET("/", showIndex.ShowIndexPage)
	router.GET("/article/view/:article_id", singleArticle.GetArticle)
}
