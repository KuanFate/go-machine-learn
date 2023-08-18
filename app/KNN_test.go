package app

import "testing"

var trainData = [][]float64{}
var targetData = []float64{18, 90}
var k = 5

func Test_knn(t *testing.T) {
	KNN(trainData, targetData, k)
}
