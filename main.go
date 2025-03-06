package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("html-template/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "Base", gin.H{"Content": "Hello World"})
	})

}
