package player

import (
	"testing"

	"github.com/Nathan-Dunne/GoLayer/element"
	"github.com/stretchr/testify/assert"
	"github.com/veandco/go-sdl2/sdl"
)

func TestPlayerCanMoveRightWhenRightKeyPress(t *testing.T) {

	player, teardownTestCase := setupPlayer(t)
	defer teardownTestCase(t)

	oldX := player.Position.X

	keyboardMover := player.GetComponent(&keyboardMover{}).(*keyboardMover)
	element.Delta = 1

	keys := sdl.GetKeyboardState()
	keys[sdl.SCANCODE_RIGHT] = 1
	keyboardMover.OnUpdate()

	t.Run("TestPlayerCanMoveUpWhenUpKeyPress", func(t *testing.T) {

		expected := oldX + keyboardMover.speed*element.Delta
		actual := player.Position.X
		assert.Equal(t, expected, actual)
	})
}
func setKeyUpdate(key int, keyboardMover *keyboardMover) {
	keys := sdl.GetKeyboardState()
	keys[key] = 1
	keyboardMover.OnUpdate()
}

func TestPlayerCanMoveLeftWhenLeftKeyPress(t *testing.T) {

	player, teardownTestCase := setupPlayer(t)
	defer teardownTestCase(t)

	oldX := player.Position.X

	keyboardMover := player.GetComponent(&keyboardMover{}).(*keyboardMover)
	setKeyUpdate(sdl.SCANCODE_LEFT, keyboardMover)

	t.Run("TestPlayerCanMoveUpWhenUpKeyPress", func(t *testing.T) {

		expected := oldX - keyboardMover.speed*element.Delta
		actual := player.Position.X
		assert.Equal(t, expected, actual)
	})
}

func TestPlayerCanMoveUpWhenUpKeyPress(t *testing.T) {

	player, teardownTestCase := setupPlayer(t)
	defer teardownTestCase(t)

	oldY := player.Position.Y

	keyboardMover := player.GetComponent(&keyboardMover{}).(*keyboardMover)
	setKeyUpdate(sdl.SCANCODE_UP, keyboardMover)

	t.Run("TestPlayerCanMoveUpWhenUpKeyPress", func(t *testing.T) {

		expected := oldY - keyboardMover.speed*element.Delta
		actual := player.Position.Y
		assert.Equal(t, expected, actual)
	})
}

func TestPlayerCanMoveDownWhenDownKeyPress(t *testing.T) {

	player, teardownTestCase := setupPlayer(t)
	defer teardownTestCase(t)

	oldY := player.Position.Y

	keyboardMover := player.GetComponent(&keyboardMover{}).(*keyboardMover)
	setKeyUpdate(sdl.SCANCODE_DOWN, keyboardMover)

	t.Run("TestPlayerCanMoveUpWhenUpKeyPress", func(t *testing.T) {

		expected := oldY + keyboardMover.speed*element.Delta
		actual := player.Position.Y
		assert.Equal(t, expected, actual)
	})
}
