package main

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

func main() {
	var GD_X []float64
	var GD_Y []float64
	x := 0.0
	alpha := 1.0
	f_change := f1(x)
	f_current := f_change
	GD_X = append(GD_X, x)
	GD_Y = append(GD_Y, f_current)
	iter_num := 0

	for f_change > 1e-10 && iter_num < 50 {
		iter_num++
		x = x - alpha*h1(x)
		tmp := f1(x)
		f_change = math.Abs(f_current - tmp)
		f_current = tmp
		GD_X = append(GD_X, x)
		GD_Y = append(GD_Y, f_current)
	}

	fmt.Printf("final result: (%.5f, %.5f)\n", x, f_current)
	fmt.Printf("numbers of iteration: %d\n", iter_num)
	fmt.Println(GD_X)
}
