package cellular

import (
	"image"

	"github.com/s0rg/grid"
)

// Moore returns slice of image.Point with Moore neighborhood diffs (all around).
func Moore() []image.Point {
	return grid.Points(grid.DirectionsALL...)
}

// VonNeumann returns slice of image.Point with von Neumann neighborhood (cross neighborhood).
func VonNeumann() []image.Point {
	return grid.Points(grid.DirectionsCardinal...)
}

// VonNeumannExtended returns slice of image.Point with von Neumann range-2 neighborhood.
func VonNeumannExtended() []image.Point {
	const two = 2

	return append(
		VonNeumann(),
		[]image.Point{
			{X: 0, Y: -two},
			{X: two, Y: 0},
			{X: 0, Y: two},
			{X: -two, Y: 0},
		}...,
	)
}
