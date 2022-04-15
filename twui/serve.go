package twui

import (
	"html/template"
	"net/http"

	"github.com/kyoto-framework/kyoto"
	"github.com/kyoto-framework/kyoto/actions"
	"github.com/kyoto-framework/kyoto/render"
	"github.com/kyoto-framework/kyoto/smode"
)

func serve(page smode.Page) {
	http.HandleFunc("/", render.PageHandler(
		smode.Adapt(page),
	))
	http.HandleFunc("/internal/actions/", actions.Handler(func(c *kyoto.Core) *template.Template {
		return template.Must(template.New("Actions").Funcs(render.FuncMap(c)).ParseGlob("*.html"))
	}))

	smode.Register(
		&AppUINavSidebar{},
		&AppUIOverlayModal{},
		&AppUIOverlaySideover{},
		&MarketingSectionHero{},
		&MarketingSectionPricing{},
	)

	println("Listening on http://localhost:25025")
	http.ListenAndServe("localhost:25025", nil)
}
