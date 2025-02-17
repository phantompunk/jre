package cmd

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phantompunk/jre.rest/assets"
	"github.com/spf13/cobra"
)

var docsCmd = &cobra.Command{
	Use: "docs",
	Run: serveDocs,
}

func serveDocs(cmd *cobra.Command, args []string) {
	router := gin.Default()
	router.StaticFS("static", http.FS(assets.AssetsFS))
	router.SetHTMLTemplate(template.Must(template.ParseFS(assets.TemplateFS, "templates/*")))

	router.GET("/", pageDocs)
	router.Run()
}

func pageDocs(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "docs.html", "")
}

func init() {
	rootCmd.AddCommand(docsCmd)
}
