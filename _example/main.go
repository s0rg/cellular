package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/s0rg/cellular"
)

func main() {
	const (
		W, H   = 20, 20
		chance = 0.3
	)

	a := cellular.New[bool](W, H, func(v bool) bool {
		return v
	})

	// prefill
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for x := 0; x < W; x++ {
		for y := 0; y < H; y++ {
			if r.Float64() < chance {
				a.Set(x, y, true)
			}
		}
	}

	const (
		// default Conway's Game of Life rules
		neighboursMin = 2
		neighboursMax = 3
	)

	// define evolution rule
	rule := func(n int, v bool) (rv bool) {
		switch v {
		case true:
			switch {
			case n < neighboursMin, n > neighboursMax: // under / over population
				rv = false
			default:
				rv = v
			}
		default:
			if n == neighboursMax { // reproduction
				rv = true
			}
		}

		return rv
	}

	// cache neighbours coords for later use
	neighbours := cellular.Moore()

	// while anybody alive...
	for i := 0; a.AliveCount() > 0; i++ {
		fmt.Printf("--- step: %d\n\n", i)
		fmt.Print(a)
		fmt.Println("-----------")

		a.Evolve(neighbours, rule)

		time.Sleep(time.Second / 2)
	}
}
