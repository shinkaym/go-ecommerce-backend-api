package basic

import (
	"testing"
)

// go test -coverprofile=coverage.out
// go tool cover -html=coverage.out -o coverage.html

func TestAddOne(t *testing.T) {
	var (
		input  = 1
		output = 2
	)

	actual := AddOne(1)
	if actual != output {
		t.Errorf("AddOne(%d), input %d, actual = %d", input, output, actual)
	}
}

func TestAddOne2(t *testing.T) {
	var (
		input  = 1
		output = 2
	)

	actual := AddOne2(1)
	if actual != output {
		t.Errorf("AddOne(%d), input %d, actual = %d", input, output, actual)
	}
}
