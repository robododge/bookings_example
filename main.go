package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/robododge/bookings-example/template"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("html-template/*")
	r.Static("/static", "./static")

	helloComp := template.Hello("Tyler")

	r.GET("/", func(c *gin.Context) {
		htmlHello, err := templ.ToGoHTML(c.Request.Context(), helloComp)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "Base", gin.H{"Content": err.Error()})
		}
		c.HTML(http.StatusOK, "Base", gin.H{"Content": htmlHello})
	})

	r.Run()

}
