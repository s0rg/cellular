package cellular

import (
	"fmt"
	"image"
	"strings"

	"github.com/s0rg/grid"
)

// Rule evolution step rule, takes number of alive neigboirs.
type Rule[T any] func(int, T) T

// Alive checks is cell (with certain value) alive.
type Alive[T any] func(T) bool

// Automata holds cells state.
type Automata[T any] struct {
	state *grid.Map[T]
	alive Alive[T]
}

// New creates cellular automata with empty state.
func New[T any](w, h int, alive Alive[T]) *Automata[T] {
	return &Automata[T]{
		alive: alive,
		state: grid.New[T](image.Rect(0, 0, w, h)),
	}
}

// Set sets cell value in given coordinates.
func (a *Automata[T]) Set(x, y int, v T) (ok bool) {
	return a.state.Set(image.Pt(x, y), v)
}

// Get returns cell value from given coordinates.
func (a *Automata[T]) Get(x, y int) (v T, ok bool) {
	return a.state.Get(image.Pt(x, y))
}

// Evolve does single evolution step, with given Rule and looking for neighbours in given directions.
func (a *Automata[T]) Evolve(
	dirs []image.Point,
	rule Rule[T],
) {
	clone := grid.New[T](a.state.Rectangle())

	a.state.Iter(func(p image.Point, v T) (next bool) {
		var aliveNeighbours int

		a.state.Neighbours(p, dirs, func(_ image.Point, v T) (next bool) {
			if a.alive(v) {
				aliveNeighbours++
			}

			return true
		})

		clone.Set(p, rule(aliveNeighbours, v))

		return true
	})

	a.state = clone
}

// AliveCount returns total alive cells count.
func (a *Automata[T]) AliveCount() (rv int) {
	a.state.Iter(func(_ image.Point, v T) (next bool) {
		if a.alive(v) {
			rv++
		}

		return true
	})

	return rv
}

// Iter iterates cells.
func (a *Automata[T]) Iter(it func(x, y int, alive bool)) {
	a.state.Iter(func(p image.Point, v T) (next bool) {
		it(p.X, p.Y, a.alive(v))

		return true
	})
}

// String returns representation of cells state.
func (a *Automata[T]) String() string {
	w, h := a.state.Bounds()

	cells := make([][]rune, h)

	for i := 0; i < h; i++ {
		cells[i] = make([]rune, w)
	}

	a.Iter(func(x, y int, alive bool) {
		var r = ' '

		if alive {
			r = '@'
		}

		cells[y][x] = r
	})

	var sb strings.Builder

	for y := 0; y < h; y++ {
		fmt.Fprintln(&sb, string(cells[y]))
	}

	return sb.String()
}
