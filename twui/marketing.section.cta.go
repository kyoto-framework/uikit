package twui

import "html/template"

type MarketingSectionCTA struct {
	Background string
	Content    MarketingSectionCTAContent
	CTA        template.HTML
}

type MarketingSectionCTAContent struct {
	Title         MarketingSectionCTAContentText
	Subtitle      MarketingSectionCTAContentText
	Description   MarketingSectionCTAContentText
	LinksCentered bool
	Links         []MarketingSectionCTAContentLink
}

type MarketingSectionCTAContentText struct {
	Text     string
	Centered bool
	Color    string
}

type MarketingSectionCTAContentLink struct {
	Href    string
	Primary bool
	Text    string
}
