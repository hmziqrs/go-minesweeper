package main

import (
	"github.com/rivo/tview"
)

func RenderGrid(r int, c int, table *tview.Table) {
	for ri := 0; ri <= r; ri++ {
		for ci := 0; ci <= c; ci++ {
			table.SetCell(ri, ci, tview.NewTableCell(" X ").SetMaxWidth(10))
		}
	}
}
