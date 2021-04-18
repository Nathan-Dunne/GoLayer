package drawing

import (
	"fmt"
	"io/ioutil"
	"path"
	"time"

	"github.com/Nathan-Dunne/GoLayer/element"
	"github.com/veandco/go-sdl2/sdl"
)

type Animator struct {
	container       *element.Element
	sequences       map[string]*Sequence
	Current         string
	lastFrameChange time.Time
	Finished        bool
}

func NewAnimator(container *element.Element, sequences map[string]*Sequence, defaultSequence string) *Animator {
	var an Animator
	an.container = container
	an.sequences = sequences
	an.Current = defaultSequence
	an.lastFrameChange = time.Now()

	return &an
}

func (an *Animator) OnUpdate() error {
	sequence := an.sequences[an.Current]
	frameInterval := float64(time.Second) / sequence.sampleRate

	if time.Since(an.lastFrameChange) >= time.Duration(frameInterval) {
		an.Finished = sequence.nextFrame()
		an.lastFrameChange = time.Now()
	}

	return nil
}

func (an *Animator) OnDraw(renderer *sdl.Renderer) error {
	tex := an.sequences[an.Current].texture()

	return DrawTexture(tex, an.container.Position, an.container.Rotation, renderer)
}

func (an *Animator) OnCollision(other *element.Element) error {
	return nil
}

func (an *Animator) SetSequence(name string) {
	an.Current = name
	an.lastFrameChange = time.Now()
}

type Sequence struct {
	textures   []*sdl.Texture
	frame      int
	sampleRate float64
	loop       bool
}

func NewSequence(filepath string, sampleRate float64, loop bool, renderer *sdl.Renderer) (*Sequence, error) {
	var seq Sequence
	files, err := ioutil.ReadDir(filepath)
	if err != nil {
		return nil, fmt.Errorf("reading dir %v: %v", filepath, err)
	}

	for _, file := range files {
		filename := path.Join(filepath, file.Name())
		tex, err := LoadTextureFromBMP(filename, renderer)
		if err != nil {
			return nil, fmt.Errorf("loading sequence frame %v", err)
		}
		seq.textures = append(seq.textures, tex)
	}

	seq.sampleRate = sampleRate
	seq.loop = loop

	return &seq, nil
}

func (seq *Sequence) texture() *sdl.Texture {
	return seq.textures[seq.frame]
}

func (seq *Sequence) nextFrame() bool {
	if seq.frame == len(seq.textures)-1 {
		if seq.loop {
			seq.frame = 0
		} else {
			return true
		}
	} else {
		seq.frame++
	}

	return false
}
