package control

import (
	"testing"
)

func AddOne(val Any) (Any, error) {
	prev := val.(int)
	return Any(prev + 1), nil
}

func TestWaterfallAddOne(t *testing.T) {

	zero := Any(0)
	four, err := Waterfall(zero, AddOne, AddOne, AddOne, AddOne)

	if err != nil {
		t.Error(err)
	} else if four.(int) != 4 {
		t.Fail()
	}

}
