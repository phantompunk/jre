package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Quote struct {
	ID    string
	Quote string
}

func quote(ctx *gin.Context) {
	res := &Quote{ID: "12345", Quote: "In ancient times, the only way you see someone like Brock Lesnar is if he came to your island in a boat, and you ran"}
	ctx.JSON(http.StatusOK, res)
}

func initRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", quote)
	return r
}

func main() {
	r := initRouter()
	r.Run(":8080")
}
