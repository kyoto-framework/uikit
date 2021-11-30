package twui

import (
	"html/template"
	"net/http"

	"github.com/kyoto-framework/kyoto"
)

func serve(page kyoto.Page) {
	ssahandler := func() http.HandlerFunc {
		return kyoto.SSAHandler(func(p kyoto.Page) *template.Template {
			return template.Must(template.New("SSA").Funcs(kyoto.TFuncMap()).ParseGlob("*.html"))
		})
	}

	http.HandleFunc("/", kyoto.PageHandler(page))
	http.HandleFunc("/SSA/", ssahandler())

	println("Listening on http://localhost:25025")
	http.ListenAndServe("localhost:25025", nil)
}
