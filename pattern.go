package main

import "fmt"

type Pattern struct {
	p    [][]uint8
	w, h int
}

// NewPattern returns pattern to be used as initial universe state.
func NewPattern(pattern [][]uint8, width, height int) (*Pattern, error) {
	if width < 1 || height < 1 {
		return nil, fmt.Errorf("pattern can't be smaller than 1x1")
	}

	if len(pattern) != height {
		return nil, fmt.Errorf("wrong pattern height provided")
	}

	// check all pattern rows to be exact width
	wr := func() int {
		for yi:=0; yi<len(pattern); yi++ {
			if len(pattern[yi]) != width {
				return yi
			}
		}

		return -1
	}()
	if wr > -1 {
		return nil, fmt.Errorf("wrong pattern width at row $%d provided", wr)
	}

	return &Pattern{
		p: pattern,
		w: width,
		h: height,
	}, nil
}
