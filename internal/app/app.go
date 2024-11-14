package app

import (
	"context"
	"embed"
	"io/fs"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/phantompunk/jre.list/sql"
)

type App struct {
	database  *sql.Database
	log *log.Logger
	port int
	assets fs.FS
	templates fs.FS
	router *gin.Engine
}

// func LoadFilesFromEmbedFS(engine *gin.Engine, embedFS embed.FS, pattern string) {
// 	templ := template.Must(template.ParseFS(embedFS, pattern))
// 	engine.SetHTMLTemplate()
// }

func New(db *sql.Database, log *log.Logger, templates, assets embed.FS) *App {
	r := gin.Default()

	return &App{
		assets: assets,
		templates: templates,
		router: r,
		database: db,
		port: 8080,
	}
}

func (a *App) Start(ctx context.Context) error {
	
	if err := a.database.Connect(); err != nil {
		log.Fatal("Not able to open database", err.Error())
		return err
	}

	// TODO load files using embed.FS 
	a.router.StaticFile("/styles.css", "./static/styles.css")
	a.router.StaticFile("/api.yaml", "./openapi.yaml")
	a.router.LoadHTMLFiles("templates/index.html", "templates/quote.html", "templates/docs.html")
	a.router.GET("/", a.pageHome)
	a.router.GET("docs", a.pageDocs)
	a.router.GET("api/text", a.pageRefresh)
	a.router.GET("api/quote", a.getRandomQuote)
	a.router.GET("api/quote/:id", a.getQuoteById)

	a.router.Run()
	return nil
}
