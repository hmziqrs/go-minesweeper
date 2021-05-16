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

// func DifficultySelectModal() *tview.Modal {
// 	// modal := tview.NewModal().SetText("Select difficulty").AddButtons(Difficulties)

// 	// return modal

// }
