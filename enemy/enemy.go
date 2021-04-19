package enemy

import (
	"fmt"

	"github.com/Nathan-Dunne/GoLayer/drawing"
	"github.com/Nathan-Dunne/GoLayer/element"
	"github.com/Nathan-Dunne/GoLayer/projectile"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	EnemySize = 105
)

func NewEnemy(renderer *sdl.Renderer,
	position element.Vector,
	idle_sprite_path string,
	destroy_sprite_path string) *element.Element {

	enemy := &element.Element{}

	enemy.Position = position
	enemy.Rotation = 180

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

	an := drawing.NewAnimator(enemy, sequences, "idle")
	enemy.AddComponent(an)

	vtb := projectile.NewEffectedByProjectile(enemy)
	enemy.AddComponent(vtb)

	col := element.Circle{
		Center: enemy.Position,
		Radius: 38,
	}
	enemy.Collisions = append(enemy.Collisions, col)

	enemy.Active = true

	return enemy
}
