package enemy

import (
	"fmt"
	"testing"

	"github.com/Nathan-Dunne/GoLayer/drawing"
	"github.com/Nathan-Dunne/GoLayer/element"
	"github.com/Nathan-Dunne/GoLayer/projectile"
	"github.com/Nathan-Dunne/GoLayer/test_utilities"
)

func setupElementWithEffectedByProjectile(t *testing.T) (*element.Element, func(t *testing.T)) {
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
	elem := &element.Element{}

	an := drawing.NewAnimator(elem, sequences, "idle")
	elem.AddComponent(an)

	effectedByProjectile := projectile.NewEffectedByProjectile(elem)
	elem.AddComponent(effectedByProjectile)

	return elem, func(t *testing.T) {
		fmt.Printf(" Teardown test case.\n")
	}
}
