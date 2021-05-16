package main

import (
	"math/rand"
	"time"
)

func GenerateGrid(r int, c int, grid [][]int, steps [][]int) {
	for ri := 0; ri <= r; ri++ {
		grid[ri] = make([]int, c+1)
		steps[ri] = make([]int, c+1)
		for ci := 0; ci <= c; ci++ {
			rand.Seed(time.Now().UnixNano() + int64(ri+ci))
			random := rand.Intn(5)
			// fmt.Println("Check %d", random)
			if random == 1 {
				grid[ri][ci] = -1
			}
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
