package player

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Nathan-Dunne/GoLayer/drawing"
	"github.com/Nathan-Dunne/GoLayer/element"
	"github.com/Nathan-Dunne/GoLayer/test_utilities"
	"github.com/stretchr/testify/assert"
	"github.com/veandco/go-sdl2/sdl"
)

func setupPlayer(t *testing.T) (*element.Element, func(t *testing.T)) {
	fmt.Printf(" Setup test case.\n")
	t.Log("setup test case")

	renderer := test_utilities.SetupRenderer()
	sprite_path := "../sprites/player.bmp"
	player := NewPlayer(renderer, sprite_path)
	element.Delta = 1

	// When finished, release and teardown.
	return player, func(t *testing.T) {
		fmt.Printf(" Teardown test case.\n")

		// Reset any pressed keys on teardown.
		keys := sdl.GetKeyboardState()
		keys[sdl.SCANCODE_LEFT] = 0
		keys[sdl.SCANCODE_RIGHT] = 0
		keys[sdl.SCANCODE_DOWN] = 0
		keys[sdl.SCANCODE_UP] = 0

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

func TestPlayerIsActiveWhenCreated(t *testing.T) {

	player, teardownTestCase := setupPlayer(t)
	defer teardownTestCase(t)

	expected := true
	actual := player.Active

	assert.Equal(t, expected, actual)
}

func TestPlayerHasPositionWhenCreated(t *testing.T) {

	player, teardownTestCase := setupPlayer(t)
	defer teardownTestCase(t)

	expected := element.Vector{
		X: screenWidth / 2.0,
		Y: screenHeight - playerSize/2.0,
	}
	actual := player.Position

	assert.Equal(t, expected, actual)
}
