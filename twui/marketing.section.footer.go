package twui

import "html/template"

type MarketingSectionFooter struct {
	Dark       bool
	Simple     MarketingSectionFooterSimple
	Main       MarketingSectionFooterMain
	Newsletter MarketingSecionsFooterNewsletter
}

type MarketingSectionFooterMain struct {
	Enabled bool

	Blocks       []MarketingSectionFooterMainBlock
	Company      MarketingSectionFooterMainCompany
	Localization MarketingSectionFooterMainLocalization
	Newsletter   MarketingSecionsFooterNewsletter
	Reverse      bool
}

type MarketingSecionsFooterNewsletter struct {
	Enabled bool

	Action string
	Method string
	Title  string
	Text   string
	Input  MarketingSectionFooterInput
	Button MarketingSectionFooterAction
}

type MarketingSectionFooterSimple struct {
	Enabled bool

	Links    []MarketingSectionFooterLink // Only if centered
	Centered bool
	Text     string
	Socials  []MarketingSectionFooterSocial
}

type MarketingSectionFooterMainBlock struct {
	Title string
	Links []MarketingSectionFooterLink
}

type MarketingSectionFooterMainCompany struct {
	Enabled bool

	Logo    template.HTML
	Text    string
	Socials []MarketingSectionFooterSocial
}

type MarketingSectionFooterMainLocalization struct {
	Enabled bool

	Title  string
	Fields []MarketingSectionFooterSelect
}

type MarketingSectionFooterLink struct {
	Text string
	Href string
}

type MarketingSectionFooterSocial struct {
	Icon template.HTML
	Href string
}

type MarketingSectionFooterSelect struct {
	Name     string
	Value    string
	Preffix  string
	Label    string
	Width    string // w-full by default, check https://tailwindcss.com/docs/width
	Options  []AppUIFormLayoutSelectOption
	Disabled bool
}

type MarketingSectionFooterSelectOptions struct {
	Label string
	Value string
}

type MarketingSectionFooterInput struct {
	Name        string
	Label       string
	Value       string
	Placeholder string
	Preffix     string
	Width       string
	Disabled    bool
}

type MarketingSectionFooterAction struct {
	Text    string
	Value   string
	Primary bool
}
