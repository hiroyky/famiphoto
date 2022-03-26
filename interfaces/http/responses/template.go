package responses

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

type htmlTemplateRenderer struct {
	templates *template.Template
}

func (t *htmlTemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func NewHtmlTemplateRenderer() *htmlTemplateRenderer {
	return &htmlTemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/pages/*.html")),
	}
}
