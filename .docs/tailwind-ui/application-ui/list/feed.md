
# Feed - `AppUIListFeed`

Basic feed component

## UI

![feed](/examples/feed.jpg)

## Notes

No notes

## Files

- [appui.list.feed.go](https://github.com/yuriizinets/kyoto-uikit/blob/master/twui/appui.list.feed.go)
- [appui.list.feed.html](https://github.com/yuriizinets/kyoto-uikit/blob/master/twui/appui.list.feed.html)

## Example

```go
func (p *PageExample) Init() {
	p.Feed = kyoto.RegC(p, &AppUIListFeed{
		Entries: []AppUIListFeedEntry{
			{
				Icon:        `<img class="h-8 w-8 rounded-full" src="https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80">`,
				Title:       "Eduardo Benz",
				Subtitle:    "Commented",
				Time:        "1d ago",
				Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
			},
			{
				Icon: `
				<div class="p-1 rounded-full bg-blue-500 text-white">
					<svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"></path></svg>
				</div>
			`,
				Title:    "Eduardo Benz",
				Time:     "2d ago",
				Subtitle: "Friend request",
			},
			{
				Icon: `
				<div class="p-1 rounded-full bg-gray-500 text-white">
					<svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"></path></svg>
				</div>
			`,
				RAW: `
				<div>
					<span class="font-semibold">Oleg Nazarov</span>
					added tags
					<span class="px-2 border rounded-full font-semibold">
						<span class="text-red-500">•</span>
						Bug
					</span>
					<span class="ml-1 px-2 border rounded-full font-semibold">
						<span class="text-indigo-500">•</span>
						Accessibility
					</span>
					<span class="ml-1">3d ago</span>
				</div>
			`,
			},
		},
	})
}
```