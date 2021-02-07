package routes

import (
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
)

func MyViewEngine() *ginview.ViewEngine {
	return ginview.New(goview.Config{
		Root:         "views",
		Extension:    ".tmpl",
		Master:       "base",
		Partials:     []string{},
		Funcs:        TemplateFunc(),
		DisableCache: false,
		Delims: goview.Delims{
			Left:  "{{",
			Right: "}}",
		},
	})
}
