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

func setupPlayertWithKeyboardMover(t *testing.T) (*element.Element, func(t *testing.T)) {
	fmt.Printf(" Setup test case.\n")
	t.Log("setup test case")

	player := &element.Element{}

	renderer := test_utilities.SetupRenderer()
	filename := "../sprites/player.bmp"
	sprite_renderer := drawing.NewSpriteRenderer(player, renderer, filename)
	player.AddComponent(sprite_renderer)

	keyboard_mover := NewKeyboardMover(player, 5)
	player.AddComponent(keyboard_mover)

	return player, func(t *testing.T) {
		fmt.Printf(" Teardown test case.\n")
	}
}

func TestPlayerCanGetKeyboardMoverComponentWhenHasKeyboardMoverComponent(t *testing.T) {

	player, teardownTestCase := setupPlayertWithKeyboardMover(t)
	defer teardownTestCase(t)

	expected := reflect.TypeOf(&KeyboardMover{})
	actual := reflect.TypeOf(player.GetComponent(&KeyboardMover{}))

	assert.Equal(t, expected, actual)
}
