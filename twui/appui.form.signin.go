package twui

import (
	"fmt"
	"html/template"
)

type AppUIFormSignIn struct {
	// Sign in config
	Logo          template.HTML
	Title         string
	ResetPassword string // url
	Socials       []AppUIFormSignInSocial
	// Form config
	Form   AppUIFormLayout   // If not configured, default will be used
	Values map[string]string // Will be mapped on init
}

type AppUIFormSignInSocial struct {
	Icon template.HTML
	Href string
}

func (c *AppUIFormSignIn) Init() {
	// Default form
	if len(c.Form.Sections) == 0 {
		c.Form.Sections = []AppUIFormLayoutSection{
			{
				Fields: []AppUIFormLayoutField{
					{
						Type:      FORM_INPUT,
						InputType: "email",
						Name:      "Email",
						Label:     "Email address",
					},
					{
						Type:      FORM_INPUT,
						InputType: "password",
						Name:      "Password",
						Label:     "Password",
					},
				},
			},
		}
		if c.ResetPassword != "" {
			c.Form.Sections[0].Fields = append(c.Form.Sections[0].Fields, AppUIFormLayoutField{
				Type: FORM_RAW,
				RAW: template.HTML(fmt.Sprintf(`
				<div class="w-full text-right">
					<a href="%s" class="text-sm text-indigo-700 font-semibold">Reset password</a>
				</div>
			`, c.ResetPassword)),
			})
		}
		c.Form.Actions = []AppUIFormLayoutAction{
			{
				Text:      "Sign In",
				Value:     "signin",
				Primary:   true,
				FullWidth: true,
			},
		}
	}
	// Default title
	if c.Title == "" {
		c.Title = "Sign in to your account"
	}
	// Map values
	// ...
}
