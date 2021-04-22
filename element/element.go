package element

import (
	"fmt"
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
)

type Vector struct {
	X, Y float64
}

var Delta float64

type CollisionCircle struct {
	Center Vector
	Radius float64
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
	Collisions []CollisionCircle
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
