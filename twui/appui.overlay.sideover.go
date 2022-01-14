package twui

import (
	"html/template"
	"math/rand"
	"strconv"

	"github.com/kyoto-framework/kyoto/smode"
)

type AppUIOverlaySideover struct {
	ID       string
	MaxWidth string // Check https://tailwindcss.com/docs/max-width
	Content  template.HTML

	InternalActivated bool
}

func (c *AppUIOverlaySideover) Init() {
	if c.ID == "" {
		c.ID = strconv.Itoa(rand.Intn(100000))
	}
}

func (c *AppUIOverlaySideover) Actions() smode.ActionMap {
	return smode.ActionMap{
		"Open": func(args ...interface{}) {
			c.InternalActivated = true
		},
		"Close": func(args ...interface{}) {
			c.InternalActivated = false
		},
		"Toggle": func(args ...interface{}) {
			c.InternalActivated = !c.InternalActivated
		},
	}
}
