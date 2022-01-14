package twui

import (
	"html/template"
	"net/http"

	"github.com/kyoto-framework/kyoto/actions"
	"github.com/kyoto-framework/kyoto/render"
	"github.com/kyoto-framework/kyoto/smode"
)

func serve(page smode.Page) {
	http.HandleFunc("/", render.PageHandler(
		smode.Adapt(page),
	))
	http.HandleFunc("/internal/actions/", actions.Handler(func() *template.Template {
		return template.Must(template.New("Actions").Funcs(render.FuncMap()).Parse("*.html"))
	}))

	println("Listening on http://localhost:25025")
	http.ListenAndServe("localhost:25025", nil)
}
