package main

import (
	"bytes"
	"fmt"
	"math"
)

// Life stores the state of current round of 'Game of Life'.
type Life struct {
	current, next *universe
	w, h          int
}

// NewLife returns current new Life game state with pattern as initial state.
func NewLife(width, height int, pattern *Pattern) (*Life, error) {
	if pattern == nil {
		return nil, fmt.Errorf("pattern can't be empty")
	}

	if width <= pattern.w  {
		return nil, fmt.Errorf("universe width must be grater than pattern width")
	}

	if height <= pattern.h {
		return nil, fmt.Errorf("universe height must be grater than pattern height")
	}

	c := newUniverse(width, height)

	c.spawnPattern(pattern, int(math.Round(float64(width)/2)), int(math.Round(float64(height)/2)))

	return &Life{
		current: c,
		next:    newUniverse(width, height),
		w:       width,
		h:       height,
	}, nil
}

// Tick updates the game by one tick.
func (l *Life) Tick() {
	// update the field in next state.
	for y := 0; y < l.h; y++ {
		for x := 0; x < l.w; x++ {
			l.next.set(x, y, l.current.Next(x, y))
		}
	}
	// Swap current and next.
	l.current, l.next = l.next, l.current
}

// String generates ASCII representation of the game.
func (l *Life) String() string {
	var buf bytes.Buffer
	for y := 0; y < l.h; y++ {
		for x := 0; x < l.w; x++ {
			b := byte('.')
			if l.current.isAlive(x, y) {
				b = '@'
			}
			buf.WriteByte(b)
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}
