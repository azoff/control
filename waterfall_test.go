package control

import (
	"testing"
)

func AddOne(val Any) (Any, error) {
	prev := val.(int)
	return prev + 1, nil
}

func TestWaterfallAddOne(t *testing.T) int {

	four, err := Waterfall(0, AddOne, AddOne, AddOne, AddOne)

	if err != nil {
		t.Error(err)
	} else if four != 4 {
		t.Fail()
	}

}
