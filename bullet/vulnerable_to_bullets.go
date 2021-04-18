package bullet

import (
	"github.com/Nathan-Dunne/GoLayer/drawing"
	"github.com/Nathan-Dunne/GoLayer/element"
	"github.com/veandco/go-sdl2/sdl"
)

type vulnerabileToBullets struct {
	container *element.Element
	animator  *drawing.Animator
}

func NewVulnerableToBullets(container *element.Element) *vulnerabileToBullets {
	return &vulnerabileToBullets{
		container: container,
		animator:  container.GetComponent(&drawing.Animator{}).(*drawing.Animator)}
}

func (vtb *vulnerabileToBullets) OnDraw(renderer *sdl.Renderer) error {
	return nil
}

func (vtb *vulnerabileToBullets) OnUpdate() error {

	if vtb.animator.Finished && vtb.animator.Current == "destroy" {
		vtb.container.Active = false
	}
	return nil
}

func (vtb *vulnerabileToBullets) OnCollision(other *element.Element) error {
	if other.Tag == "bullet" {
		vtb.animator.SetSequence("destroy")
	}

	return nil
}
