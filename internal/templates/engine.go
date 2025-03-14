package templates

import "github.com/gofiber/template/html/v2"

var TemplateEngine *html.Engine

func InitEngine() {
	TemplateEngine = html.New("internal/templates", ".html")
}
