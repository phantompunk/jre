package assets

import (
	"embed"
)

//go:embed *.css *.js *.yaml
var AssetsFS embed.FS

//go:embed templates/*.html
var TemplateFS embed.FS

// var TemplateFS fs.FS
// var AssetsFS fs.FS

func init() {
	// var err error

	// TemplateFS, err = fs.Sub(templateFS, "templates")
	// if err != nil {
	// 	panic("Failed to subtree template FS " + err.Error())
	// }
	//
	// AssetsFS, err = fs.Sub(templateFS, "static")
	// if err != nil {
	// 	panic("Failed to subtree assets FS " + err.Error())
	// }
}
