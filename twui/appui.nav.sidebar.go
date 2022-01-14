package twui

import (
	"html/template"

	"github.com/kyoto-framework/kyoto/smode"
)

type AppUINavSidebar struct {
	Logo    template.HTML
	Links   []AppUINavSidebarLink
	Search  AppUINavSidebarSearch
	Profile AppUINavSidebarProfile

	InternalHidden    bool
	InternalIconsUsed bool
}

func (c *AppUINavSidebar) Init() {
	for _, link := range c.Links {
		if link.Icon != "" {
			c.InternalIconsUsed = true
		}
	}
}

func (c *AppUINavSidebar) Actions() smode.ActionMap {
	return smode.ActionMap{
		"Toggle": func(args ...interface{}) {
			c.InternalHidden = !c.InternalHidden
		},
	}
}

type AppUINavSidebarLink struct {
	Separator bool
	Icon      template.HTML
	Text      string
	Href      string
	Counter   int
}

type AppUINavSidebarSearch struct {
	Enabled bool
	Action  string // value will be passed as Query
	Method  string
}

type AppUINavSidebarProfile struct {
	Enabled     bool
	Avatar      template.HTML
	Href        string
	Username    string
	Text        string
	SignoutHref string
}
