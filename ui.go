package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func RenderGrid(r int, c int, table *tview.Table) {
	for ri := 0; ri <= r; ri++ {
		for ci := 0; ci <= c; ci++ {
			table.SetCell(ri, ci, tview.NewTableCell(" X ").SetMaxWidth(10))
		}
	}
}

func RenderSteps(r int, c int, grid [][]int, steps [][]int, table *tview.Table) {
	for ri := 0; ri <= r; ri++ {
		for ci := 0; ci <= c; ci++ {
			value := grid[ri][ci]
			visible := steps[ri][ci]
			if visible == 1 {
				text := fmt.Sprintf(" %d ", value)
				if value == -1 {
					text = " B "
				} else if value == 0 {
					text = "   "
				}
				table.GetCell(ri, ci).SetText(text).SetTextColor(tcell.ColorRed)
			}
		}
	}
}

func DifficultySelectModal(handler func(row, column int)) tview.Primitive {
	modal := func(p tview.Primitive, width, height int) tview.Primitive {
		return tview.NewFlex().
			AddItem(nil, 0, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(nil, 0, 1, false).
				AddItem(p, height, 1, false).
				AddItem(nil, 0, 1, false), width, 1, false).
			AddItem(nil, 0, 1, false)
	}

	table := tview.NewTable().SetBorders(true)
	for i, d := range Difficulties {
		table.SetCell(i, 0, tview.NewTableCell(fmt.Sprintf("%s: %dx%d grid with %d mines", d.Label, d.R, d.C, d.B)))
	}
	table.SetSelectable(true, true).SetFixed(1, 1).SetSelectedFunc(handler)

	return modal(table, 100, 100)

}
