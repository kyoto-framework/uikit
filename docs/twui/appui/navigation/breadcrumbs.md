
# Breadcrumbs

**`AppUINavBreadcrumbs`**

Configurable breadcrumbs component

## UI

<img style="margin-top: 10px" src="/assets/examples/breadcrumbs.jpg">

## Notes

No notes

## Files

- [appui.nav.breadcrumbs.go](https://github.com/kyoto-framework/uikit/blob/master/twui/appui.nav.breadcrumbs.go)
- [appui.nav.breadcrumbs.html](https://github.com/kyoto-framework/uikit/blob/master/twui/appui.nav.breadcrumbs.html)

## Example

```go
func (p *PageExample) Init() {
	p.Breadcrumbs = kyoto.RegC(p, &AppUINavBreadcrumbs{
		Links: []AppUINavBreadcrumbsLink{
            {
                Icon: `<svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"></path></svg>`,
                Href: "/",
            },
            {
                Text: "Projects",
                Href: "#",
            },
            {
                Text: "UIKit",
                Href: "#",
            },
            {
                Text: "TWUI",
                Href: "#",
            },
        },
	})
}
```
