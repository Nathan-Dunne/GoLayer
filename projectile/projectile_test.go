package projectile

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Nathan-Dunne/GoLayer/element"
	"github.com/stretchr/testify/assert"
)

func setupElementWithProjectile(t *testing.T) (*element.Element, func(t *testing.T)) {
	fmt.Printf(" Setup test case.\n")
	t.Log("setup test case")

	elem := &element.Element{}
	projectileSpeed := 10.00
	mover := NewProjectileMover(elem, projectileSpeed)

	elem.AddComponent(mover)

	return elem, func(t *testing.T) {
		fmt.Printf(" Teardown test case.\n")
	}
}

func TestElementCanGetProjectileComponentWhenHasProjectileComponent(t *testing.T) {

	elem, teardownTestCase := setupElementWithProjectile(t)
	defer teardownTestCase(t)

	expected := reflect.TypeOf(&ProjectileMover{})
	actual := reflect.TypeOf(elem.GetComponent(&ProjectileMover{}))

	assert.Equal(t, expected, actual)
}
