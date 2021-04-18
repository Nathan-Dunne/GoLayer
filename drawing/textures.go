package drawing

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func TextureFromBMP(renderer *sdl.Renderer, filename string) *sdl.Texture {
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
