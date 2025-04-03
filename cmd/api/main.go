package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type StrResponse map[string]string

func main() {
	app := gin.Default()

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, StrResponse{"message": "pong"})
	})

	log.Fatal(app.Run(":8085"))
}
