package data

import "math/rand"

func getValues(size int) []int {
	var arr []int = make([]int, size)
	for index := 0; index < size; index++ {
		arr[index] = rand.Intn(10000)
	}

	return arr

}

func getWeights(size int) []int {
	var arr []int = make([]int, size)
	for index := 0; index < size; index++ {
		arr[index] = rand.Intn(10000)
	}

	return arr
}

func GetData(size int) map[string][]int {
	var sorts = make(map[string][]int)
	sorts["weights"] = getWeights(size)
	sorts["values"] = getValues(size)

	return sorts
}
