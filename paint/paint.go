package paint

import (
	"github.com/Nathan-Dunne/GoLayer/drawing"
	"github.com/Nathan-Dunne/GoLayer/element"
	"github.com/veandco/go-sdl2/sdl"
)

func newPaint(renderer *sdl.Renderer, paint_speed float64) *element.Element {
	paint := &element.Element{}

	spriteRenderer := drawing.NewSpriteRenderer(paint, renderer, "sprites/player_bullet.bmp")
	paint.AddComponent(spriteRenderer)

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
