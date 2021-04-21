// Package drawing implements functions for drawing on a renderer.
//
// This circle component fills in or outlines a circle based on it's container element position.
// Mostly translated to Go from https://gist.github.com/derofim/ SDL2 circle Drawing_and_Filling_Circles
package drawing

import (
	"math"

	"github.com/Nathan-Dunne/GoLayer/element"
	"github.com/veandco/go-sdl2/sdl"
)

// Circle is made up of it's container element, a center, radius, color and alpha channel.
type Circle struct {
	container  *element.Element
	Center     element.Vector
	Radius     float64
	r, g, b, a uint8
}

// NewCircle creates a circle with some radius based on the containing elements position.
func NewCircle(container *element.Element, radius float64, center element.Vector, r, g, b, a uint8) *Circle {
	return &Circle{
		container: container,
		Radius:    radius,
		Center:    center,
		r:         r,
		g:         g,
	}
}

// OnDraw satifies the element interface and draws itself on the screen with x, y, red, green, blue and alpha.
func (circle *Circle) OnDraw(renderer *sdl.Renderer) error {
	circle.fillCircle(renderer, int(circle.Center.X), int(circle.Center.Y), int(circle.Radius), 0xFF, 0x00, 0x00, 0xFF)
	return nil
}

// OnUpdate satifies the element interface.
func (circle *Circle) OnUpdate() error {
	return nil
}

// OnCollision satifies the element interface and increases the radius by 10% when collided with.
func (circle *Circle) OnCollision(other *element.Element) error {
	circle.Radius *= float64(1.1)
	return nil
}

// setPixel colors in a point or pixel on the renderer given x, y, red, green, blue and alpha.
func (circle *Circle) setPixel(renderer *sdl.Renderer, x int32, y int32, r uint8, g uint8, b uint8, a uint8) {
	renderer.SetDrawColor(r, g, b, a)
	renderer.DrawPoint(x, y)
}

// drawCircle draws an outlined circle on the renderer given x, y, red, green, blue and alpha.
func (circle *Circle) drawCircle(surface *sdl.Renderer, n_cx int, n_cy int, radius int, r uint8, g uint8, b uint8, a uint8) {
	// if the first pixel in the screen is represented by (0,0) (which is in sdl)
	// remember that the beginning of the circle is not in the middle of the pixel
	// but to the left-top from it:

	error := float64(-radius)
	x := float64(radius) - 0.5
	y := float64(0.5)
	cx := float64(n_cx) - float64(0.5)
	cy := float64(n_cy) - float64(0.5)

	for x >= y {
		circle.setPixel(surface, int32(cx+x), int32(cy+y), r, g, b, a)
		circle.setPixel(surface, int32(cx+y), int32(cy+x), r, g, b, a)

		if x != 0 {
			circle.setPixel(surface, int32(cx-x), int32(cy+y), r, g, b, a)
			circle.setPixel(surface, int32(cx+y), int32(cy-x), r, g, b, a)
		}

		if y != 0 {
			circle.setPixel(surface, int32(cx+x), int32(cy-y), r, g, b, a)
			circle.setPixel(surface, int32(cx-y), int32(cy+x), r, g, b, a)
		}

		if x != 0 && y != 0 {
			circle.setPixel(surface, int32(cx-x), int32(cy-y), r, g, b, a)
			circle.setPixel(surface, int32(cx-y), int32(cy-x), r, g, b, a)
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

// fillCircle draws a filled circle on the renderer given x, y, red, green, blue and alpha.
func (circle *Circle) fillCircle(surface *sdl.Renderer, cx int, cy int, radius int, r uint8, g uint8, b uint8, a uint8) {
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
