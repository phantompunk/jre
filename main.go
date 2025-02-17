package main

import (
	"github.com/phantompunk/jre.rest/cmd"
)

//ggo:embed assets/*.css assets/*.js assets/*.yaml
// var Assets embed.FS

//ggo:embed assets/templates/*
// var Templates embed.FS

func main() {
	cmd.Execute()
}

// func demo() {
// 	logger := log.Default()
// 	db := db.NewDatabase(db.WithBaseUrl("./database/db.db"))
//
// 	app := app.New(db, logger, Templates, Assets)
// 	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
// 	defer cancel()
//
// 	// Start the app in a goroutine
// 	errChan := make(chan error)
// 	go func() {
// 		errChan <- app.Start(ctx)
// 	}()
//
// 	// Wait for app to finish or for context cancellation
// 	select {
// 	case <-ctx.Done():
// 		logger.Println("Shutting down gracefully...")
// 	case err := <-errChan:
// 		if err != nil {
// 			logger.Fatal("App error:", err)
// 			os.Exit(1)
// 		}
// }
// }
