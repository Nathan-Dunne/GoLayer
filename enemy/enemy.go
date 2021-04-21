package enemy

import (
	"fmt"
	"math/rand"

	"github.com/Nathan-Dunne/GoLayer/drawing"
	"github.com/Nathan-Dunne/GoLayer/element"
	"github.com/Nathan-Dunne/GoLayer/projectile"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	EnemySize = 100
)

func NewEnemy(renderer *sdl.Renderer,
	position element.Vector) *element.Element {

	enemy := &element.Element{}

	enemy.Position = position
	enemy.Rotation = 180

	vtb := projectile.NewEffectedByProjectile(enemy)
	enemy.AddComponent(vtb)

	col := element.Circle{
		Center: enemy.Position,
		Radius: 20,
	}
	enemy.Collisions = append(enemy.Collisions, col)

	// It should not matter which line in a circle is checked for color.
	red := rand.Intn(255)
	green := rand.Intn(255)
	blue := rand.Intn(255)
	alpha := uint8(0xFF)

	circle := drawing.NewCircle(enemy, 20, enemy.Position, uint8(red), uint8(green), uint8(blue), alpha)
	enemy.AddComponent(circle)

	enemy.Active = true

	return enemy
}

func addAnimator(enemy *element.Element, renderer *sdl.Renderer, idle_sprite_path string, destroy_sprite_path string) *drawing.Animator {
	idleSequence, err := drawing.NewSequence(idle_sprite_path, 5, true, renderer)
	if err != nil {
		panic(fmt.Errorf("creating idle sequence %v", err))
	}
	destroySequence, err := drawing.NewSequence(destroy_sprite_path, 15, false, renderer)
	if err != nil {
		panic(fmt.Errorf("creating destroy sequence %v", err))
	}

	sequences := map[string]*drawing.Sequence{
		"idle":    idleSequence,
		"destroy": destroySequence,
	}

	return drawing.NewAnimator(enemy, sequences, "idle")
}
