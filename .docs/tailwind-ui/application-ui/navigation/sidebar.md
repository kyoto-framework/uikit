
# Sidebar - `AppUINavSidebar`

Basic sidebar component with set of "extensions", like search, profile badge, etc.

## UI

<img style="margin-top: 10px" height="500" src="/examples/sidebar.jpg">

## Notes

No notes

## Files

- [appui.nav.sidebar.go](https://github.com/yuriizinets/kyoto-uikit/blob/master/twui/appui.nav.sidebar.go)
- [appui.nav.sidebar.html](https://github.com/yuriizinets/kyoto-uikit/blob/master/twui/appui.nav.sidebar.html)

## Example

```go
func (p *PageExample) Init() {
	p.Sidebar = kyoto.RegC(p, &AppUINavSidebar{
		Logo: template.HTML(`<a href="/">
			<img src="https://raw.githubusercontent.com/yuriizinets/kyoto/master/.web/docs/.vuepress/public/kyoto.svg" class="mx-auto h-16 w-16 scale-150" />
		</a>`),
		Search: AppUINavSidebarSearch{
			Enabled: true,
		},
		Links: []AppUINavSidebarLink{
			{
				Text: "Dashboard",
				Href: "#",
				Icon: `<svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 8v8m-4-5v5m-4-2v2m-2 4h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path></svg>`,
			},
			{
				Text: "Team",
				Href: "#",
				Icon: `<svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"></path></svg>`,
			},
			{
				Text: "Projects",
				Href: "#",
				Icon: `<svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 13v-1m4 1v-3m4 3V8M8 21l4-4 4 4M3 4h18M4 4h16v12a1 1 0 01-1 1H5a1 1 0 01-1-1V4z"></path></svg>`,
			},
			{
				Text:    "Calendar",
				Href:    "#",
				Counter: 2,
				Icon:    `<svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"></path></svg>`,
			},
		},
		Profile: AppUINavSidebarProfile{
			Enabled:  true,
			Avatar:   `<img class="h-8 w-8 rounded-full" src="https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80">`,
			Username: "Tom Cook",
			Text:     "View profile",
		},
	})
}
```
