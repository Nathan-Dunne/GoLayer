package enemy

import (
	"fmt"

	"github.com/Nathan-Dunne/GoLayer/drawing"
	"github.com/Nathan-Dunne/GoLayer/element"
	"github.com/Nathan-Dunne/GoLayer/projectile"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	BasicEnemySize = 105
)

func NewBasicEnemy(renderer *sdl.Renderer, position element.Vector) *element.Element {
	basicEnemy := &element.Element{}

	basicEnemy.Position = position
	basicEnemy.Rotation = 180

	idleSequence, err := drawing.NewSequence("sprites/basic_enemy/idle", 5, true, renderer)
	if err != nil {
		panic(fmt.Errorf("creating idle sequence %v", err))
	}
	destroySequence, err := drawing.NewSequence("sprites/basic_enemy/destroy", 15, false, renderer)
	if err != nil {
		panic(fmt.Errorf("creating destroy sequence %v", err))
	}

	sequences := map[string]*drawing.Sequence{
		"idle":    idleSequence,
		"destroy": destroySequence,
	}

	an := drawing.NewAnimator(basicEnemy, sequences, "idle")
	basicEnemy.AddComponent(an)

	vtb := projectile.NewEffectedByProjectile(basicEnemy)
	basicEnemy.AddComponent(vtb)

	col := element.Circle{
		Center: basicEnemy.Position,
		Radius: 38,
	}
	basicEnemy.Collisions = append(basicEnemy.Collisions, col)

	basicEnemy.Active = true

	return basicEnemy
}
