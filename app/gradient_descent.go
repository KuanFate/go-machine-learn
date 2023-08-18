package app

import (
	"fmt"
	"math"
)

func f1(x float64) float64 {
	return 0.5 * math.Pow(x-0.25, 2)
}

func h1(x float64) float64 {
	return 0.5 * 2 * (x - 0.25)
}

// 一维梯度下降
func GradientDescent() {
	var GdX []float64
	var GdY []float64
	x := 0.0     // initial
	alpha := 1.0 // initial step size
	funcChange := f1(x)
	funcCurrent := funcChange
	GdX = append(GdX, x)
	GdY = append(GdY, funcCurrent)
	iter_num := 0

	// condition of stopping iteration
	for funcChange > 1e-10 && iter_num < 50 {
		iter_num++
		x = x - alpha*h1(x) // update x
		tmp := f1(x)
		funcChange = math.Abs(funcCurrent - tmp) // The derivative approaches 0
		funcCurrent = tmp
		GdX = append(GdX, x)
		GdY = append(GdY, funcCurrent)
	}

	fmt.Printf("final result: (%.5f, %.5f)\n", x, funcCurrent)
	fmt.Printf("numbers of iteration: %d\n", iter_num)
	fmt.Println(GdX)
	fmt.Println(GdY)
}
