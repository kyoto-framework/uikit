package twui

import (
	"html/template"

	"github.com/kyoto-framework/kyoto/smode"
)

type AppUINavNavbar struct {
	Logo          template.HTML
	Links         []AppUINavNavbarLink
	Search        AppUINavNavbarSearch
	Notifications AppUINavNavbarNotifications
	Profile       AppUINavNavbarProfile

	InternalProfileActivated    bool
	InternalMobileMenuActivated bool
}

func (c *AppUINavNavbar) Actions(p smode.Page) smode.ActionMap {
	return smode.ActionMap{
		"ToggleProfile": func(args ...interface{}) {
			if len(c.Profile.Links) != 0 {
				c.InternalProfileActivated = !c.InternalProfileActivated
			}
		},
		"ToggleMobileMenu": func(args ...interface{}) {
			c.InternalMobileMenuActivated = !c.InternalMobileMenuActivated
		},
	}
}

type AppUINavNavbarLink struct {
	Text string
	Href string
}

type AppUINavNavbarSearch struct {
	Enabled bool
	Action  string // value will be passed as Query
	Method  string
}

type AppUINavNavbarNotifications struct {
	Enabled bool
	Counter int
	Href    string
}

type AppUINavNavbarProfile struct {
	Enabled bool
	Avatar  template.HTML
	Links   []AppUINavNavbarLink
}
