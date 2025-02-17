package cmd

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phantompunk/jre.rest/assets"
	"github.com/phantompunk/jre.rest/internal/db"
	"github.com/spf13/cobra"
)

var uiCmd = &cobra.Command{
	Use: "ui",
	Run: serveUI,
}

func serveUI(cmd *cobra.Command, args []string) {
	router := gin.Default()
	router.StaticFS("static", http.FS(assets.AssetsFS))
	router.SetHTMLTemplate(template.Must(template.ParseFS(assets.TemplateFS, "templates/*")))

	router.GET("/", pageUi)
	router.GET("/api/text", pageUiNew)
	router.Run(":8081")
}

func pageUi(ctx *gin.Context) {
	quote, err := db.GetQuote()
	if err != nil {
	}
	ctx.HTML(http.StatusOK, "index.html", quote)
}

func pageUiNew(ctx *gin.Context) {
	quote, err := db.GetQuote()
	if err != nil {
	}
	ctx.HTML(http.StatusOK, "quote.html", quote)
}

func init() {
	rootCmd.AddCommand(uiCmd)
}
