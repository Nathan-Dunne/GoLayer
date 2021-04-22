package paint

import (
	"github.com/Nathan-Dunne/GoLayer/element"
	"github.com/veandco/go-sdl2/sdl"
)

type effectedByPaint struct {
	container *element.Element
}

func NewEffectedByPaint(container *element.Element) *effectedByPaint {
	return &effectedByPaint{
		container: container}
}

func (ebp *effectedByPaint) OnDraw(renderer *sdl.Renderer) error {
	return nil
}

func (ebp *effectedByPaint) OnUpdate() error {
	return nil
}

func (ebp *effectedByPaint) OnCollision(other *element.Element) error {
	if other.Tag == "paint" {

	}

	return nil
}
