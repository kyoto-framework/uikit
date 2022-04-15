package twui

import (
	"html/template"

	"github.com/kyoto-framework/kyoto/smode"
)

type AppUIHeadPage struct {
	Title       string
	Subtitle    string
	Actions     []AppUIHeadPageAction
	Meta        []AppUIHeadPageMeta
	Breadcrumbs smode.Component
	Banner      string // url
	Avatar      string // url
}

type AppUIHeadPageMeta struct {
	Icon template.HTML
	Text string
}

type AppUIHeadPageAction struct {
	Icon    template.HTML
	Text    string
	Href    string
	Action  template.JS
	Primary bool
}
