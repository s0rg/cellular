package cellular_test

import (
	"image"
	"strings"
	"testing"

	"github.com/s0rg/cellular"
)

func TestBase(t *testing.T) {
	t.Parallel()

	const (
		side = 10
		one  = 1
		out  = side + one
	)

	cells := cellular.New[bool](side, side, func(v bool) bool {
		return v
	})

	if cells.Set(out, out, true) {
		t.Fail()
	}

	if !cells.Set(one, one, true) {
		t.Fail()
	}

	if _, ok := cells.Get(out, out); ok {
		t.Fail()
	}

	if v, ok := cells.Get(one, one); !ok || !v {
		t.Fail()
	}

	var nalive int

	cells.Iter(func(_, _ int, v bool) {
		if v {
			nalive++
		}
	})

	if nalive == 0 {
		t.Fail()
	}

	if cells.AliveCount() != nalive {
		t.Fail()
	}

	if s := cells.String(); strings.Count(s, "@") != nalive {
		t.Fail()
	}
}

func TestNeighborhood(t *testing.T) {
	t.Parallel()

	var p []image.Point

	p = cellular.Moore()
	if len(p) != 8 {
		t.Fail()
	}

	p = cellular.VonNeumann()
	if len(p) != 4 {
		t.Fail()
	}

	p = cellular.VonNeumannExtended()
	if len(p) != 8 {
		t.Fail()
	}
}

func TestEvolve(t *testing.T) {
	t.Parallel()

	const (
		side       = 10
		one        = 1
		out        = side + one
		minN, maxN = 2, 3
	)

	cells := cellular.New[bool](side, side, func(v bool) bool {
		return v
	})

	dirs := cellular.Moore()

	rule := func(neighbours int, val bool) (rv bool) {
		rv = val

		if val {
			switch {
			case neighbours < minN:
				rv = false
			case neighbours > maxN:
				rv = false
			}
		} else if neighbours == maxN {
			rv = true
		}

		return rv
	}

	c := cells.AliveCount()

	cells.Evolve(dirs, rule)

	if c != cells.AliveCount() || c > 0 {
		t.Fail()
	}

	// single cell
	cells.Set(2, 2, true)
	cells.Evolve(dirs, rule)

	// will die
	if cells.AliveCount() > 0 {
		t.Fail()
	}

	// 5 cells
	cells.Set(2, 1, true)
	cells.Set(1, 2, true)
	cells.Set(3, 2, true)
	cells.Set(2, 3, true)
	cells.Set(2, 2, true)

	cells.Evolve(dirs, rule)

	// will produce 8
	if cells.AliveCount() != 8 {
		t.Fail()
	}
}
