package main

import (
	"fmt"

	"github.com/rivo/tview"
)

func main() {
	r := 9
	c := 9
	m := 10
	moves := 0
	// didWon := false

	var grid, steps [][]int

	app := tview.NewApplication()
	table := tview.NewTable().SetBorders(true)
	pages := tview.NewPages()

	table.SetSelectedFunc(func(ri int, ci int) {
		moves++
		gameover := false
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
			ShowModal(pages, "gameover", "Gameover", []string{"Quit", "Try again", "Change difficulty"}, func(i int, _ string) {
				gameover = false
				if i == 0 {
					app.Stop()
				} else if i == 1 {
					grid, steps = GenerateGrid(r, c, m)
					RenderGrid(r, c, table)
					pages.RemovePage("gameover")
				} else {
					pages.RemovePage("gameover")
					pages.SwitchToPage("menu")
				}

			})
		}
		RenderSteps(r, c, grid, steps, table)
		if !gameover && DidFinish(r, c, grid, steps) {
			title := fmt.Sprintf("Congrats! You finished the game in %d moves", moves)
			ShowModal(pages, "finish", title, []string{"Quit", "Play again", "Change difficulty"}, func(i int, _ string) {
				if i == 0 {
					app.Stop()
				} else if i == 1 {
					grid, steps = GenerateGrid(r, c, m)
					RenderGrid(r, c, table)
					pages.RemovePage("finish")
				} else {
					pages.RemovePage("finish")
					pages.SwitchToPage("menu")
				}

			})
		}

	})
	menu := DifficultySelectPage(func(row, _ int) {
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

	fmt.Println("Thank you for playing")
}
