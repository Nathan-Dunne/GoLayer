package element_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Nathan-Dunne/GoLayer/element"
	"github.com/Nathan-Dunne/GoLayer/projectile"
	"github.com/stretchr/testify/assert"
)

func setupElementWithProjectile(t *testing.T) (*element.Element, func(t *testing.T)) {
	fmt.Printf(" Setup test case.\n")
	t.Log("setup test case")

	elem := &element.Element{}
	projectileSpeed := 10.00
	mover := projectile.NewProjectileMover(elem, projectileSpeed)

	elem.AddComponent(mover)

	return elem, func(t *testing.T) {
		fmt.Printf(" Teardown test case.\n")
	}
}

func TestElementCanGetProjectileComponentWhenHasProjectileComponent(t *testing.T) {

	elem, teardownTestCase := setupElementWithProjectile(t)
	defer teardownTestCase(t)

	expected := reflect.TypeOf(&projectile.ProjectileMover{})
	actual := reflect.TypeOf(elem.GetComponent(&projectile.ProjectileMover{}))

	assert.Equal(t, expected, actual)
}
