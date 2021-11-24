package twui

import (
	"html/template"

	"github.com/kyoto-framework/kyoto"
)

type AppUIHeadPage struct {
	Title       string
	Subtitle    string
	Actions     []AppUIHeadPageAction
	Meta        []AppUIHeadPageMeta
	Breadcrumbs kyoto.Component
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
	Primary bool
}
