package showIndex

import (
	"github.com/gin-gonic/gin"
	"github.com/metalscreame/ginHelloWorld/models"
	"github.com/metalscreame/ginHelloWorld/services/render"
)

func ShowIndexPage(c *gin.Context) {
	articles := models.GetAllArticles()
	// Call the HTML method of the Context to render a template
	render.Render(c, gin.H{
		"title":   "Home Page",
		"payload": articles}, "index.html")
}