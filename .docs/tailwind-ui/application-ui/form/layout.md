
# Form Layout - `AppUIFormLayout`

Configurable Form Layout component

## UI

![form](/examples/form.jpg)

## Notes

No notes

## Files

- [appui.form.layout.go](https://github.com/yuriizinets/kyoto-uikit/blob/master/twui/appui.form.layout.go)
- [appui.form.layout.html](https://github.com/yuriizinets/kyoto-uikit/blob/master/twui/appui.form.layout.html)

## Example

```go
func (p *PageExample) Init() {
    p.FormLayout = kyoto.RegC(p, &AppUIFormLayout{
		Action: "/",
		Method: "GET",
		Sections: []AppUIFormLayoutSection{
			{
				Title:       "Profile",
				Description: "This information will be displayed publicly",
				Fields: []AppUIFormLayoutField{
					{
						Type:        FORM_INPUT,
						Name:        "Website",
						Label:       "Website",
						Preffix:     "https://",
						Placeholder: "example.com",
					},
					{
						Type:        FORM_TEXTAREA,
						Name:        "About",
						Label:       "About",
						Placeholder: "Here is some information about me ...",
						Description: "Brief description for your profile",
					},
					{
						Type:    FORM_RADIO,
						Name:    "UnderEighteen",
						Label:   "I'm under 18",
						Value:   "true",
						Checked: true,
					},
					{
						Type:  FORM_RADIO,
						Name:  "UnderEighteen",
						Label: "I'm above 18",
						Value: "false",
					},
					{
						Type:  FORM_CHECKBOX,
						Name:  "MarketingEnabled",
						Label: "Enable marketing emails",
					},
				},
			},
			{
				Title:       "Personal Information",
				Description: "Use permanent address where you can receive mail",
				Fields: []AppUIFormLayoutField{
					{
						Type:  FORM_INPUT,
						Name:  "FirstName",
						Label: "First name",
						Width: "w-1/2",
					},
					{
						Type:  FORM_INPUT,
						Name:  "LastName",
						Label: "Last name",
						Width: "w-1/2",
					},
					{
						Type:  FORM_INPUT,
						Name:  "Email",
						Label: "Email address",
					},
					{
						Type:  FORM_SELECT,
						Name:  "Country",
						Label: "Country",
						Options: []AppUIFormLayoutSelectOption{
							{Label: "United States", Value: "US"},
							{Label: "Canada", Value: "CA"},
							{Label: "Mexico", Value: "MX"},
						},
					},
					{
						Type:  FORM_INPUT,
						Name:  "City",
						Label: "City",
						Width: "w-1/3",
					},
					{
						Type:  FORM_INPUT,
						Name:  "State",
						Label: "State / Province",
						Width: "w-1/3",
					},
					{
						Type:  FORM_INPUT,
						Name:  "Zip",
						Label: "Zip / Postal code",
						Width: "w-1/3",
					},
				},
			},
		},
		Actions: []AppUIFormLayoutAction{
			{Text: "Reset", Value: "Reset"},
			{Text: "Submit", Value: "Submit", Primary: true},
		},
	})
}
```