package twui

import "html/template"

type AppUINavBreadcrumbs struct {
	Links     []AppUINavBreadcrumbsLink
	Separator string // arrow,slash
	Contained bool
	Rounded   bool
}

func (c *AppUINavBreadcrumbs) Init() {
	if c.Separator == "" {
		c.Separator = "arrow"
	}
}

type AppUINavBreadcrumbsLink struct {
	Icon template.HTML
	Text string
	Href string
}
