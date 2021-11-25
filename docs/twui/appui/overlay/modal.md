
# Modal

**`AppUIOverlayModal`**

Basic modal component with raw html content

## UI

![modalOpen](/assets/examples/modalOpen.jpg)

![modal](/assets/examples/modal.jpg)

## Notes

No notes

## Files

- [appui.overlay.modal.go](https://github.com/yuriizinets/kyoto-uikit/blob/master/twui/appui.overlay.modal.go)
- [appui.overlay.modal.html](https://github.com/yuriizinets/kyoto-uikit/blob/master/twui/appui.overlay.modal.html)

## Example

```html
<button onclick="{{ action `TestModal:Open` }}" class="py-2 px-3 bg-blue-600 text-white font-bold rounded-lg">Open</button>
```

```go
func (p *PageExample) Init() {
	p.Modal = kyoto.RegC(p, &AppUIOverlayModal{
		ID: "TestModal",
		Content: `
			<div class="p-4">
				<div class="text-xl">Test Modal</div>
				<button onclick="Action(this, 'Close')" class="mt-2 py-2 px-3 bg-blue-600 text-white font-bold rounded-lg">Close</button>
			</div>`,
		MaxWidth: "max-w-screen-sm",
	})
}
```