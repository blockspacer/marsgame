package world

import "aakimov/marsgame/go/physics"

type Object struct {
	Pos    physics.Point
	Width  int
	Height int
	Angle  float64
}
