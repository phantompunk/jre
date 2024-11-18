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

//go:embed static/*
var assets embed.FS

//go:embed templates/*
var templates embed.FS

func main() {
	log := log.Default()
	db := sql.NewDatabase(sql.WithBaseUrl("./database/db.db"))

	app := app.New(db, log, templates, assets)
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Start the app in a goroutine
	errChan := make(chan error)
	go func() {
		errChan <- app.Start(ctx)
	}()

	// Wait for app to finish or for context cancellation
	select {
	case <-ctx.Done():
		log.Println("Shutting down gracefully...")
	case err := <-errChan:
		if err != nil {
			log.Fatal("App error:", err)
			os.Exit(1)
		}
	}
}
