package enemy

import (
	"fmt"
	"testing"

	"github.com/Nathan-Dunne/GoLayer/drawing"
	"github.com/Nathan-Dunne/GoLayer/element"
	"github.com/Nathan-Dunne/GoLayer/projectile"
	"github.com/Nathan-Dunne/GoLayer/test_utilities"
	"github.com/stretchr/testify/assert"
)

func setupBasicEnemyWithEffectedByProjectile(t *testing.T) (*element.Element, func(t *testing.T)) {
	fmt.Printf(" Setup test case.\n")
	t.Log("setup test case")

	renderer := test_utilities.SetupRenderer()

	idleSequence, err := drawing.NewSequence("../sprites/basic_enemy/idle", 5, true, renderer)
	if err != nil {
		panic(fmt.Errorf("creating idle sequence %v", err))
	}
	destroySequence, err := drawing.NewSequence("../sprites/basic_enemy/destroy", 15, false, renderer)
	if err != nil {
		panic(fmt.Errorf("creating destroy sequence %v", err))
	}

	sequences := map[string]*drawing.Sequence{
		"idle":    idleSequence,
		"destroy": destroySequence,
	}
	basic_enemy := &element.Element{}

	an := drawing.NewAnimator(basic_enemy, sequences, "idle")
	basic_enemy.AddComponent(an)

	effectedByProjectile := projectile.NewEffectedByProjectile(basic_enemy)
	basic_enemy.AddComponent(effectedByProjectile)

	return basic_enemy, func(t *testing.T) {
		fmt.Printf(" Teardown test case.\n")
	}
}

func TestBasicEnemyIsNotActiveWhenCollidesWithProjectile(t *testing.T) {
	basic_enemy, teardownTestCase := setupBasicEnemyWithEffectedByProjectile(t)
	defer teardownTestCase(t)

	projectile := element.Element{}
	projectile.Tag = "projectile"

	basic_enemy.Collision(&projectile)

	assert.Equal(t, false, basic_enemy.Active)
}
