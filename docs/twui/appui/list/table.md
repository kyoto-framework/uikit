
# Table

**`AppUIListTable`**

Basic table component

## UI

![table](/assets/examples/table.jpg)

## Notes

No notes

## Files

- [appui.list.table.go](https://github.com/yuriizinets/kyoto-uikit/blob/master/twui/appui.list.table.go)
- [appui.list.table.html](https://github.com/yuriizinets/kyoto-uikit/blob/master/twui/appui.list.table.html)

## Example

```go
func (p *PageExample) Init() {
	rows := []map[string]template.HTML{}
	for i := 0; i < 100; i++ {
		rows = append(rows, map[string]template.HTML{
			"Name":   "Jane Cooper",
			"Title":  "Regional Paradigm Technician",
			"Status": `<span class="px-2 py-1 rounded-full bg-green-200 text-green-800">active</span>`,
			"Role":   "Admin",
			"Actions": `
				<div class="w-full flex justify-end gap-4">
					<a href="#" class="text-blue-500 hover:underline">Edit</a>
					<a href="#" class="text-red-500 hover:underline">Delete</a>
				</div>
			`,
		})
	}
	p.Table = kyoto.RegC(p, &AppUIListTable{
		Columns: []AppUIListTableColumn{
			{Label: "Name", Key: "Name"},
			{Label: "Title", Key: "Title"},
			{Label: "Status", Key: "Status"},
			{Label: "Role", Key: "Role"},
			{Label: "", Key: "Actions"},
		},
		Rows: rows,
	})
}
```