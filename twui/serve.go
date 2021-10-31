package twui

import (
	"html/template"
	"net/http"

	"github.com/yuriizinets/kyoto"
)

func serve(page kyoto.Page) {
	ssatemplate := func(p kyoto.Page) *template.Template {
		t := template.New("SSA").Funcs(kyoto.TFuncMap())
		t, _ = t.ParseGlob("*.html")
		t, _ = t.ParseGlob("*.gohtml")
		return t
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
