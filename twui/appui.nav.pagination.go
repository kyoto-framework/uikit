package twui

import (
	"fmt"
	"html/template"
	"strconv"
	"strings"
)

type AppUINavPagination struct {
	Page       int
	TotalPages int
	Href       string // {} will be replaced with actual page
	OnClick    string

	InternalElements []template.HTML
}

func (c *AppUINavPagination) Init() {
	// Starter
	if c.Page > 2 {
		c.InternalElements = append(c.InternalElements, c.render(1))
	}
	if c.Page > 3 {
		c.InternalElements = append(c.InternalElements, c.spacer())
	}
	// Central
	shift := 0
	if c.Page == 1 {
		shift += 1
	}
	if c.Page == c.TotalPages {
		shift -= 1
	}
	c.InternalElements = append(c.InternalElements, c.render(c.Page-1+shift))
	if c.TotalPages > 1 {
		c.InternalElements = append(c.InternalElements, c.render(c.Page+shift))
	}
	if c.TotalPages > 2 {
		c.InternalElements = append(c.InternalElements, c.render(c.Page+1+shift))
	}
	// Closer
	if c.Page < c.TotalPages-2 {
		c.InternalElements = append(c.InternalElements, c.spacer())
	}
	if c.Page < c.TotalPages-1 {
		c.InternalElements = append(c.InternalElements, c.render(c.TotalPages))
	}
}

func (c *AppUINavPagination) render(page int) template.HTML {
	currclass := ""
	if page == c.Page {
		currclass = "text-blue-600"
	}
	if c.Href != "" {
		return template.HTML(fmt.Sprintf(
			`<a href="%s" class="px-1 py-2 w-10 %s">%d</a>`,
			strings.ReplaceAll(c.Href, "{}", strconv.Itoa(page)),
			currclass,
			page,
		))
	} else {
		action := `type="submit"`
		if c.OnClick != "" {
			action = fmt.Sprintf(`onclick="%s"`, c.OnClick)
		}
		return template.HTML(fmt.Sprintf(`
			<button %s class="w-10 %s">%d</button>
		`, action, currclass, page))
	}
}

func (c *AppUINavPagination) spacer() template.HTML {
	return `<div class="px-1 py-2 w-10">...</div>`
}
