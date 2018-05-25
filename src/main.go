package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/metalscreame/ginHelloWorld/src/articles"
	"errors"
	"strconv"
)

var router *gin.Engine


var tmpArticleList []articles.Article


func main() {

	gin.SetMode(gin.ReleaseMode)
	// Set the router as the default one provided by Gin
	router = gin.Default()
	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")
	// Initialize the routes
	initializeRoutes()
	// Start serving the application
	router.Run()
}



func initializeRoutes() {
	// Handle the index route
	router.GET("/", showIndexPage)
	router.GET("/article/view/:article_id", getArticle)
	logg
}


func showIndexPage(c *gin.Context) {
	articles := articles.GetAllArticles()
	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
	)
}


func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("templates/*")
	}
	return r
}



func getArticleByID(id int) (*articles.Article, error) {
	for _, a := range articles.GetAllArticles() {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}

func getArticle(c *gin.Context) {
	// Check if the article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Check if the article exists
		if article, err := getArticleByID(articleID); err == nil {
			// Call the HTML method of the Context to render a template
			c.HTML(
				// Set the HTTP status to 200 (OK)
				http.StatusOK,
				// Use the index.html template
				"article.html",
				// Pass the data that the page uses
				gin.H{
					"title":   article.Title,
					"payload": article,
				},
			)
		} else {
			// If the article is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}