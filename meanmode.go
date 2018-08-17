/*Have the function MeanMode(arr) take the array of numbers 
stored in arr and return 1 if the mode equals the mean, 
0 if they don't equal each other (ie. [5, 3, 3, 3, 1] should 
return 1 because the mode (3) equals the mean (3)). 
The array will not be empty, will only contain positive integers, 
and will not contain more than one mode
*/


package main

import (
	"fmt"
)

func MeanMode(arr []int) int { 
	//create dictionary/hash 
	//compute mean 
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
	fmt.Println(MeanMode([]int{5,3,3,3,1})) //1
	fmt.Println(MeanMode([]int{1,2,3})) //0
}