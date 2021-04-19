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

func setupPlayer(t *testing.T) (*element.Element, func(t *testing.T)) {
	fmt.Printf(" Setup test case.\n")
	t.Log("setup test case")

	renderer := test_utilities.SetupRenderer()
	sprite_path := "../sprites/player.bmp"
	player := NewPlayer(renderer, sprite_path)

	// When finished, release and teardown.
	return player, func(t *testing.T) {
		fmt.Printf(" Teardown test case.\n")
	}
}

func TestPlayerCanGetKeyboardMoverComponentWhenCreated(t *testing.T) {

	player, teardownTestCase := setupPlayer(t)
	defer teardownTestCase(t)

	expected := reflect.TypeOf(&keyboardMover{})
	actual := reflect.TypeOf(player.GetComponent(&keyboardMover{}))

	assert.Equal(t, expected, actual)
}

func TestPlayerCanGetKeyboardShooterComponentWhenCreated(t *testing.T) {

	player, teardownTestCase := setupPlayer(t)
	defer teardownTestCase(t)

	expected := reflect.TypeOf(&keyboardShooter{})
	actual := reflect.TypeOf(player.GetComponent(&keyboardShooter{}))

	assert.Equal(t, expected, actual)
}

func TestPlayerCanGetSpriteRendererComponentWhenCreated(t *testing.T) {

	player, teardownTestCase := setupPlayer(t)
	defer teardownTestCase(t)

	expected := reflect.TypeOf(&drawing.SpriteRenderer{})
	actual := reflect.TypeOf(player.GetComponent(&drawing.SpriteRenderer{}))

	assert.Equal(t, expected, actual)
}
