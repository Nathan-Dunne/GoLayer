package drawing

import (
	"fmt"

	"github.com/Nathan-Dunne/GoLayer/element"
	"github.com/veandco/go-sdl2/sdl"
)

type SpriteRenderer struct {
	container *element.Element
	tex       *sdl.Texture

	Width, Height float64
}

// Return a POINTER to a spriteRenderer. Pointers to components and elements. Passing these around and need to make sure always refering to the same one.
func NewSpriteRenderer(container *element.Element, renderer *sdl.Renderer, filename string) *SpriteRenderer {
	tex, err := LoadTextureFromBMP(filename, renderer)
	if err != nil {
		fmt.Errorf("loading textures: %v", err)
	}

	_, _, width, height, err := tex.Query()
	if err != nil {
		panic(fmt.Errorf("querying texture: %v", err))
	}
	return &SpriteRenderer{
		container: container,
		tex:       tex,
		Width:     float64(width),
		Height:    float64(height),
	}
}

func (sr *SpriteRenderer) OnDraw(renderer *sdl.Renderer) error {
	return DrawTexture(sr.tex, sr.container.Position, sr.container.Rotation, renderer)
}

func (sr *SpriteRenderer) OnUpdate() error {
	return nil
}

func (mover *SpriteRenderer) OnCollision(other *element.Element) error {
	return nil
}
