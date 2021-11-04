package twui

type MarketingSectionCTA struct {
	Background string
	Content    MarketingSectionCTAContent
	Image      string // src
	Links      []MarketingSectionCTALink
}

type MarketingSectionCTAContent struct {
	Title         MarketingSectionCTAContentText
	Subtitle      MarketingSectionCTAContentText
	Description   MarketingSectionCTAContentText
	Centered      bool
	LinksCentered bool
	Links         []MarketingSectionCTALink
}

type MarketingSectionCTAContentText struct {
	Text     string
	Centered bool
	Color    string
}

type MarketingSectionCTALink struct {
	Href    string
	Primary bool
	Text    string
}
