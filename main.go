package main

import (
	"fmt"
	"math"
	"time"

	"github.com/Nathan-Dunne/GoLayer/collisions"
	"github.com/Nathan-Dunne/GoLayer/element"
	"github.com/Nathan-Dunne/GoLayer/enemy"
	"github.com/Nathan-Dunne/GoLayer/player"
	"github.com/Nathan-Dunne/GoLayer/projectile"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth          = 600
	screenHeight         = 800
	TargetTicksPerSecond = 60
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("Initialising SDL:", err)
		return
	}

	window, err := sdl.CreateWindow(
		"Gaming in Go", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight, sdl.WINDOW_OPENGL)
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

	idle_sprite_path := "sprites/enemy/idle"
	destroy_sprite_path := "sprites/enemy/destroy"

	for i := 0; i < 0; i++ {
		for j := 0; j < 0; j++ {
			x := (float64(i)/5)*screenWidth + (enemy.EnemySize / 2)
			y := float64(j)*enemy.EnemySize + (enemy.EnemySize / 2)

			enemy := enemy.NewEnemy(renderer, element.Vector{X: x, Y: y}, idle_sprite_path, destroy_sprite_path)
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
		fill_circle(renderer, 200, 200, 50, 0xFF, 0x00, 0x00, 0xFF)
		fill_circle(renderer, 300, 200, 50, 0x00, 0xFF, 0x00, 0xFF)
		fill_circle(renderer, 400, 200, 50, 0x00, 0x00, 0xFF, 0xFF)

		fill_circle(renderer, 300, 600, 50, 0x00, 0x00, 0x00, 0xFF)

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

// Source: https://gist.github.com/derofim/912cfc9161269336f722 SDL2 circle Drawing_and_Filling_Circles
// Translated into Go.
func set_pixel(renderer *sdl.Renderer, x int32, y int32, r uint8, g uint8, b uint8, a uint8) {
	renderer.SetDrawColor(r, g, b, a)
	renderer.DrawPoint(x, y)
}

// Source: https://gist.github.com/derofim/912cfc9161269336f722 SDL2 circle Drawing_and_Filling_Circles
// Translated into Go.
func draw_circle(surface *sdl.Renderer, n_cx int, n_cy int, radius int, r uint8, g uint8, b uint8, a uint8) {
	// if the first pixel in the screen is represented by (0,0) (which is in sdl)
	// remember that the beginning of the circle is not in the middle of the pixel
	// but to the left-top from it:

	error := float64(-radius)
	x := float64(radius) - 0.5
	y := float64(0.5)
	cx := float64(n_cx) - float64(0.5)
	cy := float64(n_cy) - float64(0.5)

	for x >= y {
		set_pixel(surface, int32(cx+x), int32(cy+y), r, g, b, a)
		set_pixel(surface, int32(cx+y), int32(cy+x), r, g, b, a)

		if x != 0 {
			set_pixel(surface, int32(cx-x), int32(cy+y), r, g, b, a)
			set_pixel(surface, int32(cx+y), int32(cy-x), r, g, b, a)
		}

		if y != 0 {
			set_pixel(surface, int32(cx+x), int32(cy-y), r, g, b, a)
			set_pixel(surface, int32(cx-y), int32(cy+x), r, g, b, a)
		}

		if x != 0 && y != 0 {
			set_pixel(surface, int32(cx-x), int32(cy-y), r, g, b, a)
			set_pixel(surface, int32(cx-y), int32(cy-x), r, g, b, a)
		}

		error += y
		y++
		error += y

		if error >= 0 {
			x--
			error -= x
			error -= x
		}

	}
}

// Source: https://gist.github.com/derofim/912cfc9161269336f722 SDL2 circle Drawing_and_Filling_Circles
// Translated into Go.
func fill_circle(surface *sdl.Renderer, cx int, cy int, radius int, r uint8, g uint8, b uint8, a uint8) {
	dy := float64(1)

	for dy <= float64(radius) {
		// This loop is unrolled a bit, only iterating through half of the
		// height of the circle.  The result is used to draw a scan line and
		// its mirror image below it.

		// We are using half of the width of the circle because we are provided
		// with a center and we need left/right coordinates.

		dx := float64(math.Floor(math.Sqrt(((float64(2.0) * float64(radius) * dy) - (dy * dy)))))
		surface.SetDrawColor(r, g, b, a)

		surface.DrawLine(int32(cx-int(dx)), int32(cy+int(dy)-radius), int32(cx+int(dx)), int32(cy+int(int(dy)-radius)))
		surface.DrawLine(int32(cx-int(dx)), int32(cy-int(dy)+radius), int32(cx+int(dx)), int32(cy-int(dy)+radius))

		dy += 1.0
	}
}
