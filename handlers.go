package main

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/robododge/bookings-example/template/view"
	"github.com/robododge/bookings-example/util"
	"github.com/robododge/bookings-example/util/gintempl"

	"github.com/gin-gonic/gin"
)

func RootHandler(c *gin.Context) {

	locale := c.GetHeader("Accept-Language")

	savingsLost := util.SavingsLost()
	dashComp := view.Dashboard(locale, util.MonthsForYear("2019"), util.RoomNightsPerMonth(),
		savingsLost)
	dashHTML, err := templ.ToGoHTML(c.Request.Context(), dashComp)
	if err != nil {
		c.HTML(http.StatusOK, "Base", gin.H{"errMsg": "OOPS! unknown error"})
	}
	debugHTML, err := templ.ToGoHTML(c.Request.Context(), view.Debugg())
	if err != nil {
		c.HTML(http.StatusOK, "Base", gin.H{"errMsg": "OOPS! unknown error"})
	}

	c.HTML(http.StatusOK, "Base", gin.H{
		"Debug":   debugHTML,
		"Content": dashHTML,
	})
}

func ChartRefresh(c *gin.Context) {
	chartID := c.Param("chartID")
	log.Printf("Refreshing chart %s", chartID)

	if chartID == "lost-savings" {
		lostSavings := util.SavingsLost()
		lostSavingsComp := view.Chart("lost-savings", util.HotelNamesOnly(lostSavings), util.SavingsOnly(lostSavings), true)

		ginSavingsLost := gintempl.New(c.Request.Context(), http.StatusOK, lostSavingsComp)
		c.Render(http.StatusOK, ginSavingsLost)
	}

}
