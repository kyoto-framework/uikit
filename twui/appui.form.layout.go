package twui

import "html/template"

const (
	FORM_INPUT    = "input"
	FORM_SELECT   = "select"
	FORM_TEXTAREA = "textarea"
	FORM_CHECKBOX = "checkbox"
	FORM_RADIO    = "radio"
	FORM_RAW      = "raw"
)

type AppUIFormLayout struct {
	Action   string
	Method   string
	Sections []AppUIFormLayoutSection
	Actions  []AppUIFormLayoutAction
}

type AppUIFormLayoutSection struct {
	Title       string
	Description string
	Fields      []AppUIFormLayoutField
}

type AppUIFormLayoutAction struct {
	Icon    template.HTML
	Text    string
	Value   string
	Primary bool
}

type AppUIFormLayoutField struct {
	Type        string
	InputType   string
	Name        string
	Value       string
	Label       string
	Preffix     string
	Description string
	Placeholder string
	Width       string // w-full by default, check https://tailwindcss.com/docs/width
	Options     []AppUIFormLayoutSelectOption
	Disabled    bool
	Checked     bool
	RAW         template.HTML
}

type AppUIFormLayoutSelectOption struct {
	Label string
	Value string
}
