package projectile

import (
	"github.com/Nathan-Dunne/GoLayer/element"
	"github.com/veandco/go-sdl2/sdl"
)

type effectedByProjetile struct {
	container *element.Element
}

func NewEffectedByProjectile(container *element.Element) *effectedByProjetile {
	return &effectedByProjetile{
		container: container}
}

func (ebp *effectedByProjetile) OnDraw(renderer *sdl.Renderer) error {
	return nil
}

func (ebp *effectedByProjetile) OnUpdate() error {
	return nil
}

func (ebp *effectedByProjetile) OnCollision(other *element.Element) error {
	if other.Tag == "projectile" {

	}

	return nil
}
