package element

import (
	"fmt"
	"math"
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
)

type Vector struct {
	X, Y float64
}

var Delta float64

type Circle struct {
	Center Vector
	Radius float64
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

// Anything that plays the part of a component needs to at least provide the methods defined in this interface.
// Something that is generally useful for most components and is expected in all of them.
type component interface {
	OnUpdate() error
	OnDraw(renderer *sdl.Renderer) error
	OnCollision(other *Element) error
} // Need to satisfy this interface.

type Element struct {
	Position   Vector
	Rotation   float64
	Active     bool // If the element is active
	Tag        string
	Collisions []Circle
	components []component
}

func (elem *Element) Draw(renderer *sdl.Renderer) error {

	for _, comp := range elem.components {
		err := comp.OnDraw(renderer)
		if err != nil {
			return nil
		}
	}

	return nil
}

func (elem *Element) Update() error {

	for _, comp := range elem.components {
		err := comp.OnUpdate()
		if err != nil {
			return nil
		}
	}

	return nil
}

func (elem *Element) Collision(other *Element) error {
	for _, comp := range elem.components {
		err := comp.OnCollision(other)
		if err != nil {
			return err
		}
	}

	return nil
}

func (elem *Element) AddComponent(new component) {
	// No real reason to want an element to have multiple components of the same type, each component type provides its own unique functionality.
	for _, existing := range elem.components {
		// Reflection allows us to look at the language itself and extract information
		if reflect.TypeOf(new) == reflect.TypeOf(existing) {
			panic(fmt.Sprintf("attempt to add new component with existing type %v", reflect.TypeOf(new)))
		}
	}

	elem.components = append(elem.components, new)
}

func (elem *Element) GetComponent(withType component) component {
	typ := reflect.TypeOf(withType)

	for _, comp := range elem.components {
		if reflect.TypeOf(comp) == typ {
			return comp
		}
	}

	panic(fmt.Sprintf("no component with type %v", reflect.TypeOf(withType)))
}

var Elements []*Element
