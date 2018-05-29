package render

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Content-type") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}
