package services

import (
	"github.com/gin-gonic/gin"
	"github.com/metalscreame/ginHelloWorld/services/singleArticle"
	"github.com/metalscreame/ginHelloWorld/services/showIndex"
)

func InitializeRoutes(router *gin.Engine) {
	// Handle the index route
	router.GET("/", showIndex.ShowIndexPage)
	router.GET("/article/view/:article_id", singleArticle.GetArticle)
}
