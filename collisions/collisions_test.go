package collisions

import (
	"testing"

	"github.com/Nathan-Dunne/GoLayer/element"
	"github.com/stretchr/testify/assert"
)

func TestCollisionCollidesGivenTwoPerfectlyOverlappingCircles(t *testing.T) {

	circle_one := element.CollisionCircle{
		Center: element.Vector{X: 0, Y: 0},
		Radius: 10,
	}

	circle_two := element.CollisionCircle{
		Center: element.Vector{X: 0, Y: 0},
		Radius: 10,
	}

	expected := true
	actual := collides(circle_one, circle_two)

	assert.Equal(t, expected, actual)

}

func TestCollisionCollidesGivenTwoNonPerfectlyOverlappingCircles(t *testing.T) {

	circle_one := element.CollisionCircle{
		Center: element.Vector{X: 5, Y: 5},
		Radius: 10,
	}

	circle_two := element.CollisionCircle{
		Center: element.Vector{X: 0, Y: 0},
		Radius: 10,
	}

	expected := true
	actual := collides(circle_one, circle_two)

	assert.Equal(t, expected, actual)

}

func TestCollisionDoesNotCollideGivenTwoNonOverlappingCircles(t *testing.T) {

	circle_one := element.CollisionCircle{
		Center: element.Vector{X: 100, Y: 100},
		Radius: 10,
	}

	circle_two := element.CollisionCircle{
		Center: element.Vector{X: 115, Y: 115},
		Radius: 10,
	}

	expected := false
	actual := collides(circle_one, circle_two)

	assert.Equal(t, expected, actual)

}
