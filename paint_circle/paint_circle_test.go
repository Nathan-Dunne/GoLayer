// Package paint_circle implements functions for creating a circle to be painted in, as an element.
//
// This test file runs unit tests around using the components that can enemies have.
package paint_circle

import (
	"fmt"
	"testing"

	"github.com/Nathan-Dunne/GoLayer/element"
	"github.com/Nathan-Dunne/GoLayer/test_utilities"
	"github.com/stretchr/testify/assert"
)

// setupPaintCircle creates a paint circle with the setup needed to exist with teardown.
func setupPaintCircle(t *testing.T) (*element.Element, func(t *testing.T)) {
	fmt.Printf(" Setup test case.\n")
	t.Log("setup test case")

	renderer := test_utilities.SetupRenderer()
	paint_circle := NewPaintCircle(renderer, element.Vector{X: 0, Y: 0})

	// When finished, release and teardown.
	return paint_circle, func(t *testing.T) {
		fmt.Printf(" Teardown test case.\n")
	}
}

// TestPaintCircleIsActiveWhenCreated creates a  paint circle and tests if it is found active upon creation.
func TestPaintCircleIsActiveWhenCreated(t *testing.T) {

	paint_circle, teardownTestCase := setupPaintCircle(t)
	defer teardownTestCase(t)

	assert.Equal(t, true, paint_circle.Active)
}
