
# Navbar - `AppUINavNavbar`

Basic navbar component with set of "extensions", like search, notifications badge, profile badge, etc.

## UI

![navbar](/examples/navbar.jpg)

## Notes

No notes

## Files

- [appui.nav.navbar.go](https://github.com/kyoto-framework/uikit/blob/master/twui/appui.nav.navbar.go)
- [appui.nav.navbar.html](https://github.com/kyoto-framework/uikit/blob/master/twui/appui.nav.navbar.html)

## Example

```go
func (p *PageExample) Init() {
    p.Navbar = kyoto.RegC(p, &twui.AppUINavNavbar{
        Logo: template.HTML(`<a href="/" class="text-white text-2xl font-bold">TWUI</a>`),
        Links: []twui.AppUINavNavbarLink{
            {Text: "Dashboard", Href: "#"},
            {Text: "Team", Href: "#"},
            {Text: "Projects", Href: "#"},
            {Text: "Calendar", Href: "#"},
        },
        Search: twui.AppUINavNavbarSearch{
            Enabled: true,
            Href:    "/",
        },
        Notifications: twui.AppUINavNavbarNotifications{
            Enabled: true,
            Href:    "#",
            Counter: 4,
        },
        Profile: twui.AppUINavNavbarProfile{
            Enabled: true,
            Avatar:  `<img class="h-8 w-8 rounded-full" src="https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80">`,
            Links: []twui.AppUINavNavbarLink{
                {Text: "Your Profile", Href: "#"},
                {Text: "Settings", Href: "#"},
                {Text: "Sign Out", Href: "#"},
            },
        },
    })
}
```
