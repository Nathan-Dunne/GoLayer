// Package enemy implements functions for creating an enemy, as an element.
//
// This test file runs unit tests around using the components that can enemies have.
package enemy

import (
	"fmt"
	"testing"

	"github.com/Nathan-Dunne/GoLayer/drawing"
	"github.com/Nathan-Dunne/GoLayer/element"
	"github.com/Nathan-Dunne/GoLayer/test_utilities"
	"github.com/stretchr/testify/assert"
)

// setupEnemy creates an enemy with the setup needed to exist with setup and teardown.
func setupEnemy(t *testing.T) (*element.Element, func(t *testing.T)) {
	fmt.Printf(" Setup test case.\n")
	t.Log("setup test case")

	renderer := test_utilities.SetupRenderer()
	idle_sprite_path := "../sprites/enemy/idle"
	destroy_sprite_path := "../sprites/enemy/destroy"
	enemy := NewEnemy(renderer, element.Vector{X: 0, Y: 0}, idle_sprite_path, destroy_sprite_path)

	// When finished, release and teardown.
	return enemy, func(t *testing.T) {
		fmt.Printf(" Teardown test case.\n")
	}
}

// TestEnemySequenceDestroyWhenCollidesWithProjectile creates a projectile element and
// collides it with an enemy, which has the effectedByProjectiles component,
// and tests if its animation sequence has been set to destroy.
func TestEnemySequenceDestroyWhenCollidesWithProjectile(t *testing.T) {
	enemy, teardownTestCase := setupEnemy(t)
	defer teardownTestCase(t)

	projectile := element.Element{}
	projectile.Tag = "projectile"

	enemy.Collision(&projectile)

	animator := enemy.GetComponent(&drawing.Animator{}).(*drawing.Animator)

	assert.Equal(t, "destroy", animator.Current)
}

// TestEnemyIsActiveWhenCreated creates an enemy and tests if it is found active upon creation.
func TestEnemyIsActiveWhenCreated(t *testing.T) {

	// Require a renderer and animations when creating an enemy to satisfy spriteRenderer and animator components.
	renderer := test_utilities.SetupRenderer()
	idle_sprite_path := "../sprites/enemy/idle"
	destroy_sprite_path := "../sprites/enemy/destroy"
	enemy := NewEnemy(renderer, element.Vector{X: 0, Y: 0}, idle_sprite_path, destroy_sprite_path)

	assert.Equal(t, true, enemy.Active)
}