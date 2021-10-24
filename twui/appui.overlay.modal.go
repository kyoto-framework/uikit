package twui

import (
	"html/template"
	"math/rand"
	"strconv"

	"github.com/yuriizinets/kyoto"
)

type AppUIOverlayModal struct {
	ID       string
	MaxWidth string // Check https://tailwindcss.com/docs/max-width
	Content  template.HTML

	InternalActivated bool
}

func (c *AppUIOverlayModal) Init() {
	if c.ID == "" {
		c.ID = strconv.Itoa(rand.Intn(100000))
	}
}

func (c *AppUIOverlayModal) Actions() kyoto.ActionMap {
	return kyoto.ActionMap{
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
