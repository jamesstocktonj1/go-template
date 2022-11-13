package main

import (
	"testing"
)


func TestMyFunction(t *testing.T) {
	testVector := []int{1, 2, 3}

	answer := MyFunction(testVector[0], testVector[1], testVector[2])
	expectedAnswer := (testVector[0] * testVector[1]) + testVector[2]
	
	if answer != expectedAnswer {
		t.Errorf("exptected function to return %d but got %d", expectedAnswer, answer)
	}
}