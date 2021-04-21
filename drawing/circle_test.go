// Package drawing implements functions for drawing on a renderer.
//
// This circle_test runs unit tests around drawing circles and checking their status.
package drawing

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/Nathan-Dunne/GoLayer/element"
	"github.com/Nathan-Dunne/GoLayer/test_utilities"
	"github.com/stretchr/testify/assert"
)

func setupCircle(t *testing.T, r, g, b uint8) (*Circle, func(t *testing.T)) {
	fmt.Printf(" Setup test case.\n")
	holder := &element.Element{}

	radius := 10.00
	center := element.Vector{X: 0, Y: 0}

	alpha := uint8(0xFF)
	circle := NewCircle(holder, radius, center, r, g, b, alpha)

	// When finished, release and teardown.
	return circle, func(t *testing.T) {
		fmt.Printf(" Teardown test case.\n")
	}
}

// getLine retrieves a random line that was used to make up the circle for the purpose of finding circle color.
func getLine(circle *Circle) line {
	rand.Seed(time.Now().UnixNano()) // It should not matter which line in a circle is checked for color.
	return circle.lines[rand.Intn(len(circle.lines)-1)]
}

// Tests if a there is a red line in a red filled in circle.
func TestCircleLineIsRedWhenCreateFilledRedCircle(t *testing.T) {

	red := uint8(0xFF)
	blue := uint8(0x00)
	green := uint8(0x00)
	circle, teardownTestCase := setupCircle(t, red, blue, green)
	defer teardownTestCase(t)

	renderer := test_utilities.SetupRenderer()
	circle.OnDraw(renderer)

	line := getLine(circle)
	expected := red
	actual := line.red
	assert.Equal(t, expected, actual)
}

// Tests if a there is a green line in a red filled in circle.
func TestCircleLineIsGreenWhenCreateFilledGreenCircle(t *testing.T) {

	red := uint8(0x00)
	green := uint8(0xFF)
	blue := uint8(0x00)
	circle, teardownTestCase := setupCircle(t, red, green, blue)
	defer teardownTestCase(t)

	renderer := test_utilities.SetupRenderer()
	circle.OnDraw(renderer)

	line := getLine(circle)
	expected := green
	actual := line.green
	assert.Equal(t, expected, actual)
}

// Tests if a there is a blue line in a red filled in circle.
func TestCircleLineIsBlueWhenCreateFilledBlueCircle(t *testing.T) {
	red := uint8(0x00)
	green := uint8(0x00)
	blue := uint8(0xFF)
	circle, teardownTestCase := setupCircle(t, red, green, blue)
	defer teardownTestCase(t)
	renderer := test_utilities.SetupRenderer()

	circle.OnDraw(renderer)

	line := getLine(circle)
	expected := blue
	actual := line.blue
	assert.Equal(t, expected, actual)
}
