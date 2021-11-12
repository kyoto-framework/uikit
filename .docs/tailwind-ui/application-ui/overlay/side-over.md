
# Side-Over - `AppUIOverlaySideover`

Basic side-over component with raw html content

## UI

![sideoverOpen](/examples/sideoverOpen.jpg)

![sideover](/examples/sideover.jpg)

## Notes

No notes

## Files

- [appui.overlay.sideover.go](https://github.com/yuriizinets/kyoto-uikit/blob/master/twui/appui.overlay.sideover.go)
- [appui.overlay.sideover.html](https://github.com/yuriizinets/kyoto-uikit/blob/master/twui/appui.overlay.sideover.html)

## Example
```html
<button onclick="{{ action `TestSideover:Open` }}" class="py-2 px-3 bg-blue-600 text-white font-bold rounded-lg">Open</button>
```

```go
func (p *PageExample) Init() {
	p.Sideover = kyoto.RegC(p, &AppUIOverlaySideover{
		ID: "TestSideover",
		Content: `
			<div class="p-4">
				<div class="text-xl">Test Side-over</div>
				<button onclick="Action(this, 'Close')" class="mt-2 py-2 px-3 bg-blue-600 text-white font-bold rounded-lg">Close</button>
			</div>`,
		MaxWidth: "max-w-md",
	})
}
```