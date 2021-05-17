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
	table.SetSelectable(true, true).Select(0, 0).SetFixed(1, 1)
}

func RenderSteps(r int, c int, grid [][]int, steps [][]int, table *tview.Table) {
	for ri := 0; ri <= r; ri++ {
		for ci := 0; ci <= c; ci++ {
			value := grid[ri][ci]
			visible := steps[ri][ci]
			if visible == 1 {
				text := fmt.Sprintf(" %d ", value)
				cell := table.GetCell(ri, ci)
				cell.SetTextColor(tcell.ColorBlue)
				if value == -1 {
					cell.SetTextColor(tcell.ColorRed)
					text = " B "
				} else if value == 0 {
					text = "   "
				}
				cell.SetText(text)
			}
		}
	}
}

func DifficultySelectPage(handler func(row, column int)) tview.Primitive {
	// modal := func(p tview.Primitive, width, height int) tview.Primitive {
	// 	return tview.NewFlex().
	// 		AddItem(nil, 0, 1, false).
	// 		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
	// 			AddItem(nil, 0, 1, false).
	// 			AddItem(p, height, 1, false).
	// 			AddItem(nil, 0, 1, false), width, 1, false).
	// 		AddItem(nil, 0, 1, false)
	// }

	table := tview.NewTable().SetBorders(true)
	for i, d := range Difficulties {
		table.SetCell(i, 0, tview.NewTableCell(fmt.Sprintf("%s: %dx%d grid with %d mines", d.Label, d.R, d.C, d.M)))
	}
	table.SetSelectable(true, true).SetFixed(1, 1).SetSelectedFunc(handler)

	// return modal(table, 100, 100)
	return table

}

// ShowModal : Convenience function to create a modal.
func ShowModal(pages *tview.Pages, label, text string, buttons []string, f func(buttonIndex int, buttonLabel string)) {
	m := tview.NewModal()
	// Set modal attributes
	m.SetText(text).
		AddButtons(buttons).
		SetDoneFunc(f).
		SetFocus(0).
		SetBackgroundColor(tcell.ColorDarkSlateGrey)

	pages.AddPage(label, m, true, false)
	pages.ShowPage(label)
}
