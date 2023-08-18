package app

import (
	"fmt"
	"math"
)

type fun func(float642 float64) float64

// one dimension function
func f1(x float64) float64 {
	return 0.5 * math.Pow(x-0.25, 2)
}

// derived function
func h1(x float64) float64 {
	return 0.5 * 2 * (x - 0.25)
}

// 一维梯度下降
func GradientDescent(x, stepSize float64, f1, h1 fun) {
	var GdX []float64
	var GdY []float64

	funcChange := f1(x)
	funcCurrent := funcChange
	GdX = append(GdX, x)
	GdY = append(GdY, funcCurrent)
	iterNum := 0

	// condition of stopping iteration
	for funcChange > 1e-10 && iterNum < 50 {
		iterNum++
		x = x - stepSize*h1(x) // update x
		tmp := f1(x)
		funcChange = math.Abs(funcCurrent - tmp) // The derivative approaches 0
		funcCurrent = tmp
		GdX = append(GdX, x)
		GdY = append(GdY, funcCurrent)
	}

	fmt.Printf("final result: (%.5f, %.5f)\n", x, funcCurrent)
	fmt.Printf("numbers of iteration: %d\n", iterNum)
	fmt.Println(GdX)
	fmt.Println(GdY)
}
