package twui

import "html/template"

type AppUIListFeed struct {
	Entries []AppUIListFeedEntry

	InternalLastIndex int
}

type AppUIListFeedEntry struct {
	Icon        template.HTML
	Title       string
	Subtitle    string
	Description string
	Time        string
	RAW         template.HTML
}

func (c *AppUIListFeed) Init() {
	c.InternalLastIndex = len(c.Entries) - 1
}
