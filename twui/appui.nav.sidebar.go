package twui

import (
	"html/template"

	"github.com/yuriizinets/kyoto"
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

func (c *AppUINavSidebar) Actions() kyoto.ActionMap {
	return kyoto.ActionMap{
		"Toggle": func(args ...interface{}) {
			c.InternalHidden = !c.InternalHidden
		},
	}
}

type AppUINavSidebarLink struct {
	Icon    template.HTML
	Text    string
	Href    string
	Counter int
}

type AppUINavSidebarSearch struct {
	Enabled bool
	Href    string // value will be passed as ?query={}
}

type AppUINavSidebarProfileLink struct {
	Text string
	Href string
}

type AppUINavSidebarProfile struct {
	Enabled  bool
	Avatar   template.HTML
	Href     string
	Username string
	Text     string
}
