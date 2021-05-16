package main

import (
	"fmt"

	"github.com/rivo/tview"
)

func main() {
	r := 9
	c := 9
	m := 10
	gameover := false
	moves := 0
	// didWon := false

	var grid, steps [][]int

	app := tview.NewApplication()
	table := tview.NewTable().SetBorders(true)

	table.Select(0, 0).SetFixed(1, 1).SetSelectedFunc(func(ri int, ci int) {
		moves++
		key := grid[ri][ci]
		if key == -1 {
			gameover = true
		} else if key > 0 {
			steps[ri][ci] = 1
		} else {
			CalculatePaths(ri, ci, r, c, grid, steps)
		}
		if gameover {
			ShowAllMines(r, c, grid, steps)
		}
		RenderSteps(r, c, grid, steps, table)
	})
	pages := tview.NewPages()
	menu := DifficultySelectModal(func(row, _ int) {
		d := Difficulties[row]
		r = d.R
		c = d.C
		m = d.M
		grid, steps = GenerateGrid(r, c, m)
		RenderGrid(r, c, table)
		pages.SwitchToPage("game")
	})
	pages.AddPage("game", table, true, false)
	pages.AddPage("menu", menu, true, true)

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

	fmt.Println("Helloww World")
}
