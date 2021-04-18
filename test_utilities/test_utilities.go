package test_utilities

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth          = 600
	screenHeight         = 800
	TargetTicksPerSecond = 60
)

func SetupRenderer() *sdl.Renderer {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("Initialising SDL:", err)
	}

	window, err := sdl.CreateWindow(
		"Gaming in Go", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight, sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("Initiasling window:", err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Initialising Renderer:", err)
	}

	return renderer
}
