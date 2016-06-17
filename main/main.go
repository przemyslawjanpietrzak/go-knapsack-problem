package main

import (
	"aisd-backpack/data"
	"aisd-backpack/dynamic"
	"aisd-backpack/recursive"
	"encoding/json"
	"fmt"
	"time"
)

type Result struct {
	Time       float64
	FinalValue int
	Result     []int
}

func showOptimalSolution(values []int, weights []int, m [][]int, W int, n int) (int, []int) {
	finalValue := 0
	result := []int{}
	for W > 0 && n > 0 {
		if m[n][W] != m[n-1][W] {
			result = append(result, n)
			W = W - weights[n]
			finalValue += values[n]
			n--
		} else {
			n--
		}
	}
	return finalValue, result
}

func main() {

	var testsCount int = 5
	var testsValues []int = []int{10000, 20000, 30000, 40000, 50000, 60000, 70000, 80000, 90000, 100000}
	var results = make([]Result, len(testsValues))
	var index int

	for _, testValue := range testsValues {
		var timeValue float64
		for i := 0; i < testsCount; i++ {

			var data = data.GetData(testValue)
			values := data["values"]
			weights := data["weights"]
			W := 10
			start := time.Now()

			recursive.KnapsackRecursive(weights, values, len(weights)-1, W)
			resultDynamic := dynamic.KnapsackDynamic(weights, values, len(weights)-1, W)
			showOptimalSolution(values, weights, resultDynamic, W, len(weights)-1)

			elapsed := time.Since(start)
			timeValue += elapsed.Seconds()
		}

		results[index] = Result{
			Time: timeValue / float64(testsCount),
		}

		index++
	}

	responseJSon, _ := json.Marshal(results)
	fmt.Println(string(responseJSon))
}
