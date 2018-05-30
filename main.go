package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/heroku/go-getting-started/services"
	"github.com/heroku/go-getting-started/database"
)

var router *gin.Engine


func main() {

	dataBase.InitDataBase()
	services.InitializeRouter()

}

