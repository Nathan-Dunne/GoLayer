package projectile

import (
	"github.com/Nathan-Dunne/GoLayer/drawing"
	"github.com/Nathan-Dunne/GoLayer/element"
	"github.com/veandco/go-sdl2/sdl"
)

type effectedByProjetile struct {
	container *element.Element
	animator  *drawing.Animator
}

func NewEffectedByProjectile(container *element.Element) *effectedByProjetile {
	return &effectedByProjetile{
		container: container,
		animator:  container.GetComponent(&drawing.Animator{}).(*drawing.Animator)}
}

func (ebp *effectedByProjetile) OnDraw(renderer *sdl.Renderer) error {
	return nil
}

func (ebp *effectedByProjetile) OnUpdate() error {

	if ebp.animator.Finished && ebp.animator.Current == "destroy" {
		ebp.container.Active = false
	}
	return nil
}

func (ebp *effectedByProjetile) OnCollision(other *element.Element) error {
	if other.Tag == "projectile" {
		ebp.animator.SetSequence("destroy")
	}

	return nil
}
