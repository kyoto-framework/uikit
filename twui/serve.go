package twui

import (
	"html/template"
	"net/http"

	"github.com/yuriizinets/kyoto"
)

func serve(page kyoto.Page) {
	ssatemplate := func(p kyoto.Page) *template.Template {
		return template.Must(template.New("SSA").Funcs(kyoto.TFuncMap()).ParseGlob("*.html"))
	}
	ssahandler := func() http.HandlerFunc {
		return func(rw http.ResponseWriter, r *http.Request) {
			kyoto.SSAHandlerFactory(ssatemplate, map[string]interface{}{
				"internal:rw": rw,
				"internal:r":  r,
			})(rw, r)
		}
	}

	http.HandleFunc("/", kyoto.PageHandler(page))
	http.HandleFunc("/SSA/", ssahandler())

	println("Listening on http://localhost:25025")
	http.ListenAndServe("localhost:25025", nil)
}
