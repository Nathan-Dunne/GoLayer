package bullet

import (
	"github.com/Nathan-Dunne/GoLayer/drawing"
	"github.com/Nathan-Dunne/GoLayer/element"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	bulletSize  = 32
	bulletSpeed = 10
)

func newBullet(renderer *sdl.Renderer) *element.Element {
	bullet := &element.Element{}

	sr := drawing.NewSpriteRenderer(bullet, renderer, "sprites/player_bullet.bmp")
	bullet.AddComponent(sr)

	mover := NewProjectileMover(bullet, bulletSpeed)
	bullet.AddComponent(mover)

	col := element.Circle{
		Center: bullet.Position,
		Radius: 8,
	}

	bullet.Collisions = append(bullet.Collisions, col)

	bullet.Tag = "bullet"
	return bullet
}

// Instead of holding values or bullets, we hold pointers to be able to pass by reference.
var BulletPool []*element.Element

func InitBulletPool(renderer *sdl.Renderer) {
	for i := 0; i < 30; i++ {
		bul := newBullet(renderer) // Make the bullet.
		element.Elements = append(element.Elements, bul)
		BulletPool = append(BulletPool, bul) // Store the pointer to the bullet.
	}
}

func BulletFromPool() (*element.Element, bool) {
	// Comma ok pattern? Operations that may or may not succeed but failure is not an error.
	for _, bullet := range BulletPool {
		if !bullet.Active {
			return bullet, true
		}

	}

	return nil, false
}
