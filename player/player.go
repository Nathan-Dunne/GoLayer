package player

import (
	"time"

	"github.com/Nathan-Dunne/GoLayer/drawing"
	"github.com/Nathan-Dunne/GoLayer/element"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed        = 0.05
	playerSize         = 105
	screenWidth        = 1920
	screenHeight       = 1080
	playerShotCooldown = time.Millisecond * 250
)

// Generic element defined by the components that are attached to it.
func NewPlayer(renderer *sdl.Renderer, sprite_path string) *element.Element {
	player := &element.Element{}

	radius := 50.00

	player.Position = element.Vector{
		X: screenWidth / 2.0,
		Y: screenHeight - radius/2.0,
	}

	red := uint8(0xFF)
	green := uint8(0xFF)
	blue := uint8(0xFF)
	alpha := uint8(0xFF)

	drawing_circle := drawing.NewCircle(player, radius, player.Position, uint8(red), uint8(green), uint8(blue), alpha)
	player.AddComponent(drawing_circle)

	mover := NewKeyboardMover(player, 5)
	player.AddComponent(mover)

	shooter := NewKeyboardShooter(player, playerShotCooldown)
	player.AddComponent(shooter)

	player.Active = true

	return player
}
