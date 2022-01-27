package main

import "math"

// universe is current two-dimensional universe of cells.
type universe struct {
	field [][]bool
	w, h  int
}

// newUniverse returns an empty universe of the specified width and height.
func newUniverse(width, height int) *universe {
	f := make([][]bool, height)
	for i := range f {
		f[i] = make([]bool, width)
	}
	return &universe{field: f, w: width, h: height}
}

// set sets the state of the specified cell.
func (f *universe) set(x, y int, b bool) {
	f.field[y][x] = b
}

// isAlive returns if the specified cell is alive.
// If the x or y coordinates are outside the universe boundaries they will be wrapped
// For example, an x value of -1 is treated as width-1.
func (f *universe) isAlive(x, y int) bool {
	x += f.w
	x %= f.w
	y += f.h
	y %= f.h
	return f.field[y][x]
}

// Next returns the state of the specified cell at the next time step
func (f *universe) Next(x, y int) bool {
	// Count how many neighbors are alive near the given cell
	aliveCount := 0
	for xi := -1; xi <= 1; xi++ {
		for yi := -1; yi <= 1; yi++ {
			if (yi != 0 || xi != 0) && f.isAlive(x+xi, y+yi) {
				aliveCount++
			}
		}
	}
	// Return next state according to the game rules:
	// Any live cell with fewer than two live neighbors dies as if caused by underpopulation
	// Any live cell with two or three live neighbors lives on to the next generation
	// Any live cell with more than three live neighbors dies, as if by overcrowding
	// Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction
	return ((aliveCount == 2 || aliveCount == 3) && f.isAlive(x, y)) || (aliveCount == 3 && !f.isAlive(x, y))
}

// spawnPattern spawns given pattern at the given coords.
func (f *universe) spawnPattern(pattern *Pattern, x, y int) {
	if pattern == nil {
		return // nothing to do
	}

	// calculate middle of the pattern
	pMiddleX := int(math.Round(float64(pattern.w) / 2))
	pMiddleY := int(math.Round(float64(pattern.h) / 2))

	// spawn the pattern
	for yi := 1; yi <= pattern.h; yi++ {
		for xi := 1; xi <= pattern.w; xi++ {
			if pattern.p[yi-1][xi-1] > 0 {
				f.set(x+(xi-pMiddleX), y+(yi-pMiddleY), true)
			}
		}
	}
}
