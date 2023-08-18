package app

import (
	"fmt"
	"math"
)

func KNN(trainData [][]float64, targetData []float64, k int) {
	listDistance := [][]float64{}
	for _, train := range trainData {
		distance := 0.0
		for i := 0; i < len(train)-1; i++ {
			distance += math.Pow(train[i]-targetData[i], 2)
		}
		distance = math.Sqrt(distance)
		listDistance = append(listDistance, []float64{distance, train[len(train)-1]})
	}
	sortByDistance(listDistance)

	weights := []float64{}
	for _, i := range listDistance[:k] {
		// Adding 0.001 to prevent division by zero
		weight := 1 / (i[0] + 0.001)
		weights = append(weights, weight)
	}
	weightsSum := sum(weights)
	result := 0.0
	for i := 0; i < k; i++ {
		weights[i] = weights[i] / weightsSum
		num := weights[i] * listDistance[i][1]
		result += num
	}
	predict := 1
	if result < 0 {
		predict = -1
	}
	fmt.Println(predict)
}

func sortByDistance(list [][]float64) {
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i][0] > list[j][0] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}

func sum(nums []float64) float64 {
	total := 0.0
	for _, num := range nums {
		total += num
	}
	return total
}
