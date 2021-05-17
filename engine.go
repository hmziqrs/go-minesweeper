package main

import (
	"math/rand"
	"time"
)

func GenerateGrid(r, c, m int) ([][]int, [][]int) {
	count := 0

	grid := make([][]int, r+1)
	steps := make([][]int, r+1)
	for ri := 0; ri <= r; ri++ {
		grid[ri] = make([]int, c+1)
		steps[ri] = make([]int, c+1)
	}

	for count < m {
		rand.Seed(time.Now().UnixNano() + int64(count))
		ri := rand.Intn(r)
		ci := rand.Intn(c)
		key := grid[ri][ci]
		if key == -1 {
			continue
		} else {
			grid[ri][ci] = -1
			count++
		}
	}

	for ri := 0; ri <= r; ri++ {
		for ci := 0; ci <= c; ci++ {
			if grid[ri][ci] == -1 {
				continue
			}
			count := 0
			for nri := -1; nri <= 1; nri++ {
				for nci := -1; nci <= 1; nci++ {
					cri := ri + nri
					cci := ci + nci
					if cri < 0 || cri > r || cci < 0 || cci > c || (nri == 0 && nci == 0) {
						continue
					}
					if grid[cri][cci] == -1 {
						count++
					}
				}
			}
			grid[ri][ci] = count
		}
	}

	return grid, steps
}

func DidFinish(r int, c int, grid [][]int, steps [][]int) bool {
	for ri := 0; ri <= r; ri++ {
		for ci := 0; ci <= c; ci++ {
			key := grid[ri][ci]
			visible := steps[ri][ci]
			if key != -1 && visible == 0 {
				return false
			}
		}
	}

	return true

}

func ShowAllMines(r int, c int, grid [][]int, steps [][]int) {
	for ri := 0; ri <= r; ri++ {
		for ci := 0; ci <= c; ci++ {
			key := grid[ri][ci]
			if key == -1 {
				steps[ri][ci] = 1
			}
		}
	}
}

func CalculatePaths(sri int, sci int, r int, c int, grid [][]int, steps [][]int) {
	for nri := -1; nri <= 1; nri++ {
		for nci := -1; nci <= 1; nci++ {
			cri := sri + nri
			cci := sci + nci
			if cri < 0 || cri > r || cci < 0 || cci > c {
				continue
			}

			key := grid[cri][cci]
			if steps[cri][cci] == 1 {
				continue
			}

			if key == -1 {
				continue
			}
			steps[cri][cci] = 1
			if key == 0 {
				CalculatePaths(cri, cci, r, c, grid, steps)
			}
		}
	}
}
