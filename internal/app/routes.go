package app

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *App) pageHome(ctx *gin.Context) {
	quote, _ := a.database.GetRandomQoute(ctx)
	ctx.HTML(http.StatusOK, "index.html", quote)
}

func (a *App) pageDocs(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "docs.html", "")
}

func (a *App) pageRefresh(ctx *gin.Context) {
	quote, err := a.database.GetRandomQoute(ctx)
	checkErr(err)

	ctx.HTML(http.StatusOK, "quote.html", quote)
}

func (a *App) getRandomQuote(ctx *gin.Context) {
	quote, err := a.database.GetRandomQoute(ctx) 
	checkErr(err)

	if quote == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No quotes found"})
	} else {
		ctx.JSON(http.StatusOK, quote)
	}
}

func (a *App) getQuoteById(ctx *gin.Context) {
	id := ctx.Param("id")
	quote, err := a.database.GetQouteById(ctx, id)
	checkErr(err)

	if quote == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Quote not found"})
	} else {
		ctx.JSON(http.StatusOK, quote)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Printf("%v",err.Error())
	}
}

