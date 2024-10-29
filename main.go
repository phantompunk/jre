package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

type Database struct {
	SqlDB *sql.DB
}

type Quote struct {
	ID    string `json:"id"`
	Quote string `json:"quote"`
}

func GetRandomQuote() (*Quote, error) {
	res, err := DB.Query("SELECT id, quote FROM quotes ORDER BY RANDOM() LIMIT 1;")
	if err != nil {
		log.Fatal("Error occurred while querying database", err.Error())
		return nil, err
	}
	defer res.Close()

	quote := &Quote{}
	for res.Next() {
		err = res.Scan(&quote.ID, &quote.Quote)

		if err != nil {
			return nil, err
		}
	}
	return quote, nil
}

func quote(ctx *gin.Context) {
	quote, err := GetRandomQuote()
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, quote)
}

func getQuoteById(ctx *gin.Context) {
	quote, err := GetRandomQuote()
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, quote)
}

func initRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", quote)
	v1 := r.Group("/v1")
	{
		v1.GET("quote", quote)
		v1.GET("quote/:id", getQuoteById)
	}
	return r
}

func initDB() error {
	db, err := sql.Open("sqlite3", "./db.db")
	if err != nil {
		log.Fatal("Not able to open database", err.Error())
		return err
	}
	DB = db
	return nil
}

func main() {
	initDB()
	r := initRouter()
	r.Run(":8080")
}
