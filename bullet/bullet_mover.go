package bullet

import (
	"math"

	"github.com/Nathan-Dunne/GoLayer/element"
	"github.com/veandco/go-sdl2/sdl"
)

type bulletMover struct {
	container *element.Element
	speed     float64
}

func newBulletMover(container *element.Element, speed float64) *bulletMover {
	return &bulletMover{
		container: container,
		speed:     speed,
	}
}

func (mover *bulletMover) OnDraw(renderer *sdl.Renderer) error {
	return nil
}

func (mover *bulletMover) OnUpdate() error {
	cont := mover.container
	cont.Position.X += mover.speed * math.Cos(cont.Rotation) * element.Delta
	cont.Position.Y += mover.speed * math.Sin(cont.Rotation) * element.Delta

	if cont.Position.X > 600 || cont.Position.X < 0 ||
		cont.Position.Y > 800 || cont.Position.Y < 0 {
		cont.Active = false
	}

	mover.container.Collisions[0].Center = cont.Position

	return nil
}

func (mover *bulletMover) OnCollision(other *element.Element) error {
	mover.container.Active = false
	return nil
}
