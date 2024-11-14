package main

import (
	"context"
	"embed"
	"log"
	"os"
	"os/signal"

	"github.com/phantompunk/jre.list/internal/app"
	"github.com/phantompunk/jre.list/sql"
)

type Quote struct {
	ID      string `json:"id"`
	Quote   string `json:"quote"`
	Speaker string `json:"speaker"`
	Source  string `json:"source"`
	Data    string `json:"date"`
}

//go:embed static/*
var assets embed.FS

//go:embed templates/*.html
var templates embed.FS

func main() {
	log := log.Default()
	db := sql.NewDatabase(sql.WithBaseUrl("./database/db.db"))

	app := app.New(db, log, templates, assets)
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := app.Start(ctx); err != nil {
		log.Fatal("Failed to start app", err.Error())
		os.Exit(1)
	}
}
