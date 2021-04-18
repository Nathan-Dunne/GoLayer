package drawing

import (
	"fmt"

	"github.com/Nathan-Dunne/GoLayer/element"
	"github.com/veandco/go-sdl2/sdl"
)

func DrawTexture(tex *sdl.Texture, position element.Vector, rotation float64, renderer *sdl.Renderer) error {

	_, _, width, height, err := tex.Query()
	if err != nil {
		panic(fmt.Errorf("querying texture: %v", err))
	}

	// Converting coordinates to top left of sprite.
	// Transform x and y to represent middle of sprite.
	position.X -= float64(width) / 2.0
	position.Y -= float64(height) / 2.0

	return renderer.CopyEx(tex,
		&sdl.Rect{X: 0, Y: 0, W: int32(width), H: int32(height)},                                 // src Rect describes what part of the given texture we want to put on screen.
		&sdl.Rect{X: int32(position.X), Y: int32(position.Y), W: int32(width), H: int32(height)}, // Need to convert from float64 to int32 because pixels.
		rotation,
		&sdl.Point{X: int32(width) / 2, Y: int32(height / 2)},
		sdl.FLIP_NONE) // dst Rect describes what area this texture is going to be applied to.
	// SDL coordinates only deal with integers. As uses pixels as unit position. There is no space between pixels to draw.
}

func LoadTextureFromBMP(filename string, renderer *sdl.Renderer) (*sdl.Texture, error) {
	img, err := sdl.LoadBMP(filename)
	if err != nil {
		return nil, fmt.Errorf("loading %v: %v", filename, err)
	}
	defer img.Free() // Make sure we clean up, might have memory leaks if not careful. After in texture we don't need img any more.

	// Load image into ram, texture into gpu ram.
	// := is not assignment, it is a short variable declaration. Assignment uses e.g. the simple = operator.
	// As its name says: it is to declare variables
	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		return nil, fmt.Errorf("creating texture from %v: %v", filename, err)
	}

	return tex, nil
}
