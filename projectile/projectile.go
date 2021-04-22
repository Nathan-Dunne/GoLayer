package projectile

import (
	"github.com/Nathan-Dunne/GoLayer/drawing"
	"github.com/Nathan-Dunne/GoLayer/element"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	projectileSize  = 32
	projectileSpeed = 10
)

func newProjectile(renderer *sdl.Renderer) *element.Element {
	projectile := &element.Element{}

	spriteRenderer := drawing.NewSpriteRenderer(projectile, renderer, "sprites/player_bullet.bmp")
	projectile.AddComponent(spriteRenderer)

	projectileMover := NewProjectileMover(projectile, projectileSpeed)
	projectile.AddComponent(projectileMover)

	col := element.CollisionCircle{
		Center: projectile.Position,
		Radius: 8,
	}

	projectile.Collisions = append(projectile.Collisions, col)

	projectile.Tag = "projectile"
	return projectile
}

// Instead of holding values or bullets, we hold pointers to be able to pass by reference.
var ProjectilePool []*element.Element

func InitProjectilePool(renderer *sdl.Renderer) {
	for i := 0; i < 30; i++ {
		projectile := newProjectile(renderer) // Make the bullet.
		element.Elements = append(element.Elements, projectile)
		ProjectilePool = append(ProjectilePool, projectile) // Store the pointer to the bullet.
	}
}

func ProjectileFromPool() (*element.Element, bool) {
	// Comma ok pattern? Operations that may or may not succeed but failure is not an error.
	for _, projectile := range ProjectilePool {
		if !projectile.Active {
			return projectile, true
		}

	}

	return nil, false
}
