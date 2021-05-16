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
			// Here we compute
			fc := ci > 0
			lc := ci < c
			fr := ri > 0
			lr := ri < r
			pci := ci - 1
			pri := ri - 1
			nci := ci + 1
			nri := ri + 1
			if fc && fr && grid[pri][pci] == -1 {
				count++
			}
			if fc && grid[ri][pci] == -1 {
				count++
			}
			if fc && lr && grid[nri][pci] == -1 {
				count++
			}
			if fr && grid[pri][ci] == -1 {
				count++
			}
			if lr && grid[nri][ci] == -1 {
				count++
			}
			if fr && lc && grid[pri][nci] == -1 {
				count++
			}
			if lc && grid[ri][nci] == -1 {
				count++
			}
			if lc && lr && grid[nri][nci] == -1 {
				count++
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
	if sri > 0 {
		for ri := sri; ri >= 0; ri-- {
			key := grid[ri][sci]
			if key == -1 {
				break
			}
			steps[ri][sci] = 1
			if key > 0 {
				break
			}

			if sci > 0 {
				for ci := sci; ci >= 0; ci-- {
					key := grid[ri][ci]
					if key == -1 {
						break
					}
					steps[ri][ci] = 1
					if key > 0 {
						break
					}
				}
			}

			if sci < c {
				for ci := sci; ci <= c; ci++ {
					key := grid[ri][ci]
					if key == -1 {
						break
					}
					steps[ri][ci] = 1
					if key > 0 {
						break
					}
				}
			}
		}
	}

	if sri < r {
		for ri := sri; ri <= r; ri++ {
			key := grid[ri][sci]
			if key == -1 {
				break
			}
			steps[ri][sci] = 1
			if key > 0 {
				break
			}

			if sci > 0 {
				for ci := sci; ci >= 0; ci-- {
					key := grid[ri][ci]
					if key == -1 {
						break
					}
					steps[ri][ci] = 1
					if key > 0 {
						break
					}
				}
			}

			if sci < c {
				for ci := sci; ci <= c; ci++ {
					key := grid[ri][ci]
					if key == -1 {
						break
					}
					steps[ri][ci] = 1
					if key > 0 {
						break
					}
				}
			}
		}
	}

	if sci > 0 {
		for ci := sci; ci >= 0; ci-- {
			key := grid[sri][ci]
			if key == -1 {
				break
			}
			steps[sri][ci] = 1
			if key > 0 {
				break
			}

			if sri > 0 {
				for ri := sri; ri >= 0; ri-- {
					key := grid[ri][ci]
					if key == -1 {
						break
					}
					steps[ri][ci] = 1
					if key > 0 {
						break
					}
				}
			}
			if sri < r {
				for ri := sri; ri <= r; ri++ {
					key := grid[ri][ci]
					if key == -1 {
						break
					}
					steps[ri][ci] = 1
					if key > 0 {
						break
					}
				}
			}
		}
	}

	if sci < c {
		for ci := sci; ci <= c; ci++ {
			key := grid[sri][ci]
			if key == -1 {
				break
			}
			steps[sri][ci] = 1
			if key > 0 {
				break
			}

			if sri > 0 {
				for ri := sri; ri >= 0; ri-- {
					key := grid[ri][ci]
					if key == -1 {
						break
					}
					steps[ri][ci] = 1
					if key > 0 {
						break
					}
				}
			}
			if sri < r {
				for ri := sri; ri <= r; ri++ {
					key := grid[ri][ci]
					if key == -1 {
						break
					}
					steps[ri][ci] = 1
					if key > 0 {
						break
					}
				}
			}
		}
	}
}
