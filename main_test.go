package main

import (
	"reflect"
	"testing"
)

func TestEqualPlayers(t *testing.T) {
	expected := Player{
		name: "Mark",
		hp:   100,
	}
	have := Player{
		name: "Mark",
		hp:   100,
	}

	// Normal Test (good for single field)
	if expected.hp != have.hp {
		t.Errorf("expected %+v but got %+v", expected, have)
	}

	// DeepEqual Test (good for multiple field at once)
	if !reflect.DeepEqual(expected, have) {
		t.Errorf("expected %+v but got %+v", expected, have)
	}

}

func TestCalculateValues(t *testing.T) {
	var (
		expected = 10
		a        = 5
		b        = 5
	)

	have := calculateValues(a, b)
	if have != expected {
		t.Errorf("expected %d but have %d", expected, have)
	}
}
