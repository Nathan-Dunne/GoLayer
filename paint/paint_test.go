package paint

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Nathan-Dunne/GoLayer/element"
	"github.com/stretchr/testify/assert"
)

func setupElementWithPaint(t *testing.T) (*element.Element, func(t *testing.T)) {
	fmt.Printf(" Setup test case.\n")
	t.Log("setup test case")

	elem := &element.Element{}
	paintSpeed := 10.00
	mover := NewPaintMover(elem, paintSpeed)

	elem.AddComponent(mover)

	return elem, func(t *testing.T) {
		fmt.Printf(" Teardown test case.\n")
	}
}

func TestElementCanGetPaintComponentWhenHasPaintComponent(t *testing.T) {

	elem, teardownTestCase := setupElementWithPaint(t)
	defer teardownTestCase(t)

	expected := reflect.TypeOf(&PaintMover{})
	actual := reflect.TypeOf(elem.GetComponent(&PaintMover{}))

	assert.Equal(t, expected, actual)
}
