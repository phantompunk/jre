package app

import (
	"context"
	"embed"
	"html/template"
	"io/fs"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phantompunk/jre.list/sql"
)

type App struct {
	database  *sql.Database
	log       *log.Logger
	port      int
	assets    fs.FS
	templates fs.FS
	router    *gin.Engine
}

func LoadFilesFromEmbedFS(engine *gin.Engine, embedFS fs.FS, pattern string) {
	templ := template.Must(template.ParseFS(embedFS, pattern))
	engine.SetHTMLTemplate(templ)
}

func LoadStaticFilesFromEmbedFS(engine *gin.Engine, embedFS fs.FS, pattern string) {
	// Serve the embedded static files
	staticServer := http.FS(embedFS)
	engine.StaticFS(pattern, staticServer)
}

func New(db *sql.Database, log *log.Logger, templates, assets embed.FS) *App {
	r := gin.Default()

	return &App{
		assets:    assets,
		templates: templates,
		router:    r,
		database:  db,
		port:      8080,
	}
}

func (a *App) Start(ctx context.Context) error {

	if err := a.database.Connect(); err != nil {
		log.Fatal("Not able to open database", err.Error())
		return err
	}

	LoadFilesFromEmbedFS(a.router, a.templates, "templates/*")
	// TODO load files using embed.FS
	// LoadStaticFilesFromEmbedFS(a.router, a.assets, "/static")
	a.router.StaticFile("/styles.css", "./static/styles.css")
	a.router.StaticFile("/api.yaml", "./static/api.yaml")
	a.router.GET("/", a.pageHome)
	a.router.GET("docs", a.pageDocs)
	a.router.GET("api/text", a.pageRefresh)
	a.router.GET("api/quote", a.getRandomQuote)
	a.router.GET("api/quote/:id", a.getQuoteById)

	a.router.Run()
	return nil
}
