package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("html-template/*")
	r.Static("/static", "./static")

	r.GET("/", RootHandler)
	r.GET("/refresh-chart/:chartID", ChartRefresh)

	// helloComp := template.Hello("Tyler")

	// r.GET("/", func(c *gin.Context) {
	// 	htmlHello, err := templ.ToGoHTML(c.Request.Context(), helloComp)
	// 	if err != nil {
	// 		c.HTML(http.StatusInternalServerError, "Base", gin.H{"Content": err.Error()})
	// 	}
	// 	c.HTML(http.StatusOK, "Base", gin.H{"Content": htmlHello})
	// })

	r.Run()

}
