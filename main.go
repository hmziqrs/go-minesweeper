package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	r := 9
	c := 9
	// moves := 0
	// finish := false

	grid := make([][]int, r+1)
	steps := make([][]int, r+1)

	GenerateGrid(r, c, grid, steps)

	app := tview.NewApplication()
	table := tview.NewTable().SetBorders(true)

	RenderGrid(r, c, table)

	table.SetSelectable(true, true)

	table.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			table.SetSelectable(true, true)
		}
	}).SetSelectedFunc(func(row int, column int) {
		table.GetCell(row, column).SetTextColor(tcell.ColorRed)
		// table.SetSelectable(false, false)
	})
	if err := app.SetRoot(table, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

	fmt.Println("Helloww World")
}
