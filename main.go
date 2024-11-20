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

//go:embed assets/*.css assets/*.js assets/*.yaml
var assets embed.FS

//go:embed assets/templates/*
var templates embed.FS

func main() {
	logger := log.Default()
	db := sql.NewDatabase(sql.WithBaseUrl("./database/db.db"))

	app := app.New(db, logger, templates, assets)
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
		logger.Println("Shutting down gracefully...")
	case err := <-errChan:	
		if err != nil {
			logger.Fatal("App error:", err)
			os.Exit(1)
		}
	}
}
