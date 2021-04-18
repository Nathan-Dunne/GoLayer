package player

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Nathan-Dunne/GoLayer/drawing"
	"github.com/Nathan-Dunne/GoLayer/element"
	"github.com/Nathan-Dunne/GoLayer/test_utilities"
	"github.com/stretchr/testify/assert"
)

func setupElementWithKeyboardMover(t *testing.T) (*element.Element, func(t *testing.T)) {
	fmt.Printf(" Setup test case.\n")
	t.Log("setup test case")

	elem := &element.Element{}

	renderer := test_utilities.SetupRenderer()
	filename := "../sprites/player.bmp"
	sprite_renderer := drawing.NewSpriteRenderer(elem, renderer, filename)
	elem.AddComponent(sprite_renderer)

	keyboard_mover := NewKeyboardMover(elem, 5)
	elem.AddComponent(keyboard_mover)

	return elem, func(t *testing.T) {
		fmt.Printf(" Teardown test case.\n")
	}
}

func TestElementCanGetKeyboardMoverComponentWhenHasKeyboardMoverComponent(t *testing.T) {

	elem, teardownTestCase := setupElementWithKeyboardMover(t)
	defer teardownTestCase(t)

	expected := reflect.TypeOf(&KeyboardMover{})
	actual := reflect.TypeOf(elem.GetComponent(&KeyboardMover{}))

	assert.Equal(t, expected, actual)
}
