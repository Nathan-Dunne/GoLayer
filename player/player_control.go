package player

import (
	"math"
	"time"

	"github.com/Nathan-Dunne/GoLayer/drawing"
	"github.com/Nathan-Dunne/GoLayer/element"
	"github.com/Nathan-Dunne/GoLayer/paint"
	"github.com/veandco/go-sdl2/sdl"
)

type keyboardMover struct {
	container *element.Element
	speed     float64

	sr *drawing.Circle
}

func NewKeyboardMover(container *element.Element, speed float64) *keyboardMover {
	return &keyboardMover{
		container: container,
		speed:     speed,
		sr:        container.GetComponent(&drawing.Circle{}).(*drawing.Circle),
	}
}

func (mover *keyboardMover) OnUpdate() error {

	// A slice of uint 8's. Each elememt in slice represents the state of a key.
	// Index array using using scan codes. Scan codes are unique indexes for every keyboard key.
	keys := sdl.GetKeyboardState()

	cont := mover.container

	// Coordinates represent center of player so we need to offset by half their size.
	// X: 0, Y:0 is top left of coordinate system, Y increases down screen.
	if keys[sdl.SCANCODE_LEFT] == 1 {
		if cont.Position.X-(mover.sr.Radius/2) > 0 {
			cont.Position.X -= mover.speed * element.Delta
		}
	}
	if keys[sdl.SCANCODE_RIGHT] == 1 {
		if cont.Position.X+(mover.sr.Radius/2) < screenWidth {
			cont.Position.X += mover.speed * element.Delta
		}
	}
	if keys[sdl.SCANCODE_UP] == 1 {
		if cont.Position.Y-(mover.sr.Radius/2) >= 0 {
			cont.Position.Y -= mover.speed * element.Delta
		}
	}
	if keys[sdl.SCANCODE_DOWN] == 1 {
		if cont.Position.Y+(mover.sr.Radius/2) <= screenHeight {
			cont.Position.Y += mover.speed * element.Delta
		}
	}

	return nil
}

func (mover *keyboardMover) OnDraw(renderer *sdl.Renderer) error {
	return nil
}

type keyboardShooter struct {
	container *element.Element
	Cooldown  time.Duration
	LastShot  time.Time
}

func NewKeyboardShooter(container *element.Element, cooldown time.Duration) *keyboardShooter {
	return &keyboardShooter{
		container: container,
		Cooldown:  cooldown,
	}

}

func (mover *keyboardShooter) OnUpdate() error {
	keys := sdl.GetKeyboardState()

	pos := mover.container.Position

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(mover.LastShot) >= mover.Cooldown {
			mover.Shoot(pos.X, pos.Y-50, 270) // Up
			mover.Shoot(pos.X, pos.Y+50, 90)  // Down

			mover.Shoot(pos.X+50, pos.Y, 0)   // Right
			mover.Shoot(pos.X-50, pos.Y, 180) // Left

			mover.LastShot = time.Now()
		}
	}

	return nil
}

func (mover *keyboardMover) OnCollision(other *element.Element) error {
	return nil
}

func (mover *keyboardShooter) OnDraw(renderer *sdl.Renderer) error {
	return nil
}

func (mover *keyboardShooter) Shoot(x, y, rotation float64) {
	if paint_projectile, ok := paint.PaintFromPool(); ok {
		paint_projectile.Active = true
		paint_projectile.Position.X = x
		paint_projectile.Position.Y = y
		paint_projectile.Rotation = rotation * (math.Pi / 180)

	}
}

func (mover *keyboardShooter) OnCollision(other *element.Element) error {
	return nil
}
