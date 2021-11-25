
# Pagination

**`AppUINavPagination`**

Basic pagination component

## UI

![pagination](/assets/examples/pagination.jpg)

## Notes

No notes

## Files

- [appui.nav.pagination.go](https://github.com/yuriizinets/kyoto-uikit/blob/master/twui/appui.nav.pagination.go)
- [appui.nav.pagination.html](https://github.com/yuriizinets/kyoto-uikit/blob/master/twui/appui.nav.pagination.html)

## Example

```go
func (p *PageExample) Init() {
	p.Pagination = kyoto.RegC(p, &AppUINavPagination{
		Page:       97,
		TotalPages: 100,
	})
}
```