package main

import (
	"database/sql"
	"fmt"
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

func main() {
	initializeDB()
	r := gin.Default()
	r.GET("quote", getRandomQuote)
	r.GET("quote/:id", getQuoteById)

	// By default serves on :8080 unless a
	// PORT environment variable is supplied
	r.Run()
}

func GetRandomQuote() (*Quote, error) {
	quote := &Quote{}
	err := DB.QueryRow("SELECT id, quote FROM quotes ORDER BY RANDOM() LIMIT 1;").Scan(&quote.ID, &quote.Quote)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Error occured while querying for random quote: %v", err.Error())
		}
		return nil, fmt.Errorf("Unknown error occured querying database: %v", err.Error())
	}
	return quote, nil
}

func GetQuoteById(id string) (*Quote, error) {
	quote := &Quote{}
	err := DB.QueryRow("SELECT id, quote FROM quotes WHERE id = ?;", id).Scan(&quote.ID, &quote.Quote)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Error occured while querying for quote id:%v, %v", id, err.Error())
		}
		return nil, fmt.Errorf("Unknown error occured querying database for id:%v, %v", id, err.Error())
	}
	return quote, nil
}

func getRandomQuote(ctx *gin.Context) {
	quote, err := GetRandomQuote()
	checkErr(err)

	if quote == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No quotes found"})
	} else {
		ctx.JSON(http.StatusOK, quote)
	}
}

func getQuoteById(ctx *gin.Context) {
	id := ctx.Param("id")
	quote, err := GetQuoteById(id)
	checkErr(err)

	if quote == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Quote not found"})
	} else {
		ctx.JSON(http.StatusOK, quote)
	}
}

func initializeDB() error {
	db, err := sql.Open("sqlite3", "./db.db")
	if err != nil {
		log.Fatal("Not able to open database", err.Error())
		return err
	}
	DB = db
	return nil
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
