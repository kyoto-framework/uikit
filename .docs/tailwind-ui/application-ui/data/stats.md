
# Stats - `AppUIDataStats`

Stats component
## UI

![stats](/examples/stats.jpg)

## Notes

No notes

## Files

- [appui.data.stats.go](https://github.com/yuriizinets/kyoto-uikit/blob/master/twui/appui.data.stats.go)
- [appui.data.stats.html](https://github.com/yuriizinets/kyoto-uikit/blob/master/twui/appui.data.stats.html)

## Example

```go
func (p *PageExample) Init() {
	p.Stats = kyoto.RegC(p, &AppUIDataStats{
		Items: []AppUIDataStatsItem{
			{
				Image: `
					<div class="p-2 bg-indigo-500 text-white rounded-lg">
						<svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path></svg>
					</div>
				`,
				Title:  "Total Subscribers",
				Value:  71897,
				Change: 122,
			},
			{
				Title:        "Avg. Open Rate",
				Value:        58.16,
				ValueSymbol:  "%",
				Change:       5.4,
				ChangeSymbol: "%",
			},
			{
				Title:        "Avg. Click Rate",
				Value:        24.57,
				ValueSymbol:  "%",
				Change:       -3.2,
				ChangeSymbol: "%",
			},
		},
	})
}
```