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
	tex := textureFromBMP(renderer, filename)

	_, _, width, height, err := tex.Query()
	if err != nil {
		panic(fmt.Errorf("querying texture: %v", err))
	}
	return &SpriteRenderer{
		container: container,
		tex:       textureFromBMP(renderer, filename),
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

func textureFromBMP(renderer *sdl.Renderer, filename string) *sdl.Texture {
	img, err := sdl.LoadBMP(filename)
	if err != nil {
		panic(fmt.Errorf("loading %v: %v", filename, err))
	}
	defer img.Free() // Make sure we clean up, might have memory leaks if not careful. After in texture we don't need img any more.

	// Load image into ram, texture into gpu ram.
	// := is not assignment, it is a short variable declaration. Assignment uses e.g. the simple = operator.
	// As its name says: it is to declare variables
	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		panic(fmt.Errorf("creating texture from %v: %v", filename, err))
	}

	return tex
}
