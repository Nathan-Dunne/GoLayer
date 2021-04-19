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
	screenWidth        = 600
	screenHeight       = 800
	playerShotCooldown = time.Millisecond * 250
)

// Generic element defined by the components that are attached to it.
func NewPlayer(renderer *sdl.Renderer, sprite_path string) *element.Element {
	player := &element.Element{}

	player.Position = element.Vector{
		X: screenWidth / 2.0,
		Y: screenHeight - playerSize/2.0,
	}

	player.Active = true
	sr := drawing.NewSpriteRenderer(player, renderer, sprite_path)
	player.AddComponent(sr)

	mover := NewKeyboardMover(player, 5)
	player.AddComponent(mover)

	shooter := NewKeyboardShooter(player, playerShotCooldown)
	player.AddComponent(shooter)

	return player
}
