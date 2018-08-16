package main

import (
	"fmt"
)

func MeanMode(arr []int) int { 
	//create dictionary/hash 
	//computer mean 
	arrMap := make(map[int]int)
	total := 0
	length := len(arr)
	for _, val := range arr {
		total += val
		arrMap[val]++
	}
	mean := float64(total/length)

	//get list of map values and find max
	mapValues := []int{}
	for key := range arrMap {
		mapValues = append(mapValues,arrMap[key])
	}		
	mode := mapValues[0]
	for _, val := range mapValues {
		if val > mode {
			mode = val
		}
	}

	//return final result, 1 or 0
	if int(mean) == mode {
		return 1
	}
	return 0

}

func main() {

	fmt.Println(MeanMode([]int{5,3,3,3,1}))
}