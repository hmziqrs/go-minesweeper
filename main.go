package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	r := 9
	c := 9
	gameover := false
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
	}).SetSelectedFunc(func(ri int, ci int) {
		key := grid[ri][ci]
		if key == -1 {
			gameover = true
		}
		if gameover {
			ShowAllMines(r, c, grid, steps)
		}
		RenderSteps(r, c, grid, steps, table)

		// 		key := grid[ri][ci]
		// char := fmt.Sprintf(" %d ", key)
		// if key == -1 {
		// 	char = " B "
		// }

		// table.GetCell(ri, ci).SetTextColor(tcell.ColorRed).SetText(char)
	})
	if err := app.SetRoot(table, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

	fmt.Println("Helloww World")
}
