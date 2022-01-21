package main

import (
	"fmt"
	"time"
)

func main() {
	// create Glider pattern
	pattern, err := NewPattern([][]uint8{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
	}, 3, 3)
	if err != nil {
		panic(err)
	}

	l, err := NewLife(25, 25, pattern)
	if err != nil {
		panic(err)
	}

	for {
		l.Tick()
		fmt.Println(l)
		time.Sleep(time.Millisecond * 500)
	}
}
