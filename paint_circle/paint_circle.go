package paint_circle

import (
	"math/rand"

	"github.com/Nathan-Dunne/GoLayer/drawing"
	"github.com/Nathan-Dunne/GoLayer/element"
	"github.com/Nathan-Dunne/GoLayer/paint"
	"github.com/veandco/go-sdl2/sdl"
)

func NewPaintCircle(renderer *sdl.Renderer,
	position element.Vector) *element.Element {

	paint_circle := &element.Element{}

	paint_circle.Position = position
	paint_circle.Rotation = 180

	vtb := paint.NewEffectedByPaint(paint_circle)
	paint_circle.AddComponent(vtb)

	collision_circle := element.CollisionCircle{
		Center: paint_circle.Position,
		Radius: 20,
	}
	paint_circle.Collisions = append(paint_circle.Collisions, collision_circle)

	// It should not matter which line in a circle is checked for color.
	red := rand.Intn(255)
	green := rand.Intn(255)
	blue := rand.Intn(255)
	alpha := uint8(0xFF)

	drawing_circle := drawing.NewCircle(paint_circle, 20, paint_circle.Position, uint8(red), uint8(green), uint8(blue), alpha)
	paint_circle.AddComponent(drawing_circle)

	paint_circle.Active = true

	return paint_circle
}
