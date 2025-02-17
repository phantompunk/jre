package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/phantompunk/jre.rest/assets"
	"github.com/phantompunk/jre.rest/internal/app"
	"github.com/phantompunk/jre.rest/internal/db"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use: "api",
	Run: serveApi,
}

func serveApi(cmd *cobra.Command, args []string) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	logger := log.Default()
	db := db.NewDatabase(db.WithBaseUrl("./database/db.db"))

	app := app.New(db, logger, assets.TemplateFS, assets.AssetsFS)

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

func init() {
	rootCmd.AddCommand(apiCmd)
}
