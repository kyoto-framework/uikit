package twui

import (
	"html/template"

	"github.com/kyoto-framework/kyoto/smode"
)

type MarketingSectionHero struct {
	BackgroundImage string
	Navbar          MarketingSectionHeroNavbar
	Body            MarketingSectionHeroBody
}

func (c *MarketingSectionHero) Actions(p smode.Page) smode.ActionMap {
	return smode.ActionMap{
		"ToggleMobileMenu": func(args ...interface{}) {
			c.Navbar.InternalMobileMenuOpened = !c.Navbar.InternalMobileMenuOpened
		},
		"ToggleDropdown": func(args ...interface{}) {
			if len(args) == 0 || args[0].(string) == c.Navbar.InternalOpenedDropdown {
				c.Navbar.InternalOpenedDropdown = ""
			} else {
				c.Navbar.InternalOpenedDropdown = args[0].(string)
			}
		},
	}
}

type MarketingSectionHeroNavbar struct {
	Enabled bool

	Centered          bool
	Logo              template.HTML
	Links             []MarketingSectionHeroNavbarLink
	AutorizationLinks []MarketingSectionHeroNavbarLink

	InternalMobileMenuOpened bool
	InternalOpenedDropdown   string
}

type MarketingSectionHeroNavbarLink struct {
	Text string
	Href string

	Links []MarketingSectionHeroNavbarLink
}

type MarketingSectionHeroBody struct {
	Title    string
	Subtitle string
	Text     string
	Centered bool
	Image    MarketingSectionHeroBodyImage
	Actions  []MarketingSectionHeroBodyActions
}

type MarketingSectionHeroBodyActions struct {
	Text    string
	Href    string
	Primary bool
}

type MarketingSectionHeroBodyImage struct {
	Src    string
	Cover  bool
	Angled bool
}
