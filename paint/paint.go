package paint

import (
	"github.com/Nathan-Dunne/GoLayer/drawing"
	"github.com/Nathan-Dunne/GoLayer/element"
	"github.com/veandco/go-sdl2/sdl"
)

func newPaint(renderer *sdl.Renderer, paint_speed float64) *element.Element {
	paint := &element.Element{}

	red := uint8(0xFF)
	green := uint8(0xFF)
	blue := uint8(0xFF)
	alpha := uint8(0xFF)

	radius := 10.00

	drawing_circle := drawing.NewCircle(paint, radius, element.Vector{X: 0, Y: 0}, uint8(red), uint8(green), uint8(blue), alpha)
	paint.AddComponent(drawing_circle)

	paint_mover := NewPaintMover(paint, paint_speed)
	paint.AddComponent(paint_mover)

	col := element.CollisionCircle{
		Center: paint.Position,
		Radius: 8,
	}

	paint.Collisions = append(paint.Collisions, col)

	paint.Tag = "paint"
	return paint
}

// Instead of holding values or bullets, we hold pointers to be able to pass by reference.
var PaintPool []*element.Element

func InitPaintPool(renderer *sdl.Renderer) {
	paint_speed := 10.00

	for i := 0; i < 30; i++ {
		paint := newPaint(renderer, paint_speed) // Make the bullet.
		element.Elements = append(element.Elements, paint)
		PaintPool = append(PaintPool, paint) // Store the pointer to the bullet.
	}
}

func PaintFromPool() (*element.Element, bool) {
	// Comma ok pattern? Operations that may or may not succeed but failure is not an error.
	for _, paint := range PaintPool {
		if !paint.Active {
			return paint, true
		}

	}

	return nil, false
}
