package twui

import "html/template"

type AppUIListTable struct {
	Columns []AppUIListTableColumn
	Rows    []map[string]template.HTML
}

type AppUIListTableColumn struct {
	Label string
	Key   string
}
