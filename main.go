package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Nathan-Dunne/GoLayer/collisions"
	"github.com/Nathan-Dunne/GoLayer/element"
	"github.com/Nathan-Dunne/GoLayer/enemy"
	"github.com/Nathan-Dunne/GoLayer/player"
	"github.com/Nathan-Dunne/GoLayer/projectile"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	ScreenWidth          = 1920
	ScreenHeight         = 1080
	TargetTicksPerSecond = 60
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("Initialising SDL:", err)
		return
	}

	window, err := sdl.CreateWindow(
		"Gaming in Go", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		ScreenWidth, ScreenHeight, sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("Initiasling window:", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Initialising Renderer:", err)
		return
	}
	defer renderer.Destroy()

	sprite_path := "sprites/player.bmp"
	firstPlayer := player.NewPlayer(renderer, sprite_path)
	element.Elements = append(element.Elements, firstPlayer)

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			x := float64(rand.Intn(ScreenWidth - 1))
			y := float64(rand.Intn(ScreenHeight - 1))

			enemy := enemy.NewEnemy(renderer, element.Vector{X: x, Y: y})
			element.Elements = append(element.Elements, enemy)
		}
	}

	projectile.InitProjectilePool(renderer)

	for {

		frameStartTime := time.Now()
		// Whenever anything happens SDL puts it on a queue.
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			// Events in SDL are always pointer objects.
			case *sdl.QuitEvent:
				return
			}
		}

		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		for _, elem := range element.Elements {
			if elem.Active {
				err = elem.Update()
				if err != nil {
					fmt.Println("updating element:", err)
				}

				err = elem.Draw(renderer)
				if err != nil {
					fmt.Println("drawing element:", err)
				}

			}
		}

		if err := collisions.CheckCollisions(); err != nil {
			fmt.Println("checkking collisions", err)
			return
		}

		// x and y value for top left coord. Width and Height goes out through those points.
		// draw rect on img to copy
		// sprites sheets
		// & Ambersand syntax, literal rect that will be a pointer
		// dst rect allows change to image

		renderer.Present()

		element.Delta = time.Since(frameStartTime).Seconds() * TargetTicksPerSecond
	}
}
