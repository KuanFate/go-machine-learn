package app

import (
	"testing"
)

var (
	x        float64 = 0.0 // initial
	stepSize float64 = 1.0 // initial step size
)

func TestGradientDescent(t *testing.T) {
	GradientDescent(x, stepSize, f1, h1)
}
