package paint

import (
	"math"

	"github.com/Nathan-Dunne/GoLayer/element"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 1920
	screenHeight = 1080
)

type PaintMover struct {
	container *element.Element
	speed     float64
}

func NewPaintMover(container *element.Element, speed float64) *PaintMover {
	return &PaintMover{
		container: container,
		speed:     speed,
	}
}

func (mover *PaintMover) OnDraw(renderer *sdl.Renderer) error {
	return nil
}

func (mover *PaintMover) OnUpdate() error {
	cont := mover.container
	cont.Position.X += mover.speed * math.Cos(cont.Rotation) * element.Delta
	cont.Position.Y += mover.speed * math.Sin(cont.Rotation) * element.Delta

	if cont.Position.X > screenWidth || cont.Position.X < 0 ||
		cont.Position.Y > screenHeight || cont.Position.Y < 0 {
		cont.Active = false
	}

	mover.container.Collisions[0].Center = cont.Position

	return nil
}

func (mover *PaintMover) OnCollision(other *element.Element) error {
	mover.container.Active = false
	return nil
}
