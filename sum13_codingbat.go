/*
Return the sum of the numbers in the array, returning 0 for an empty array. 
Except the number 13 is very unlucky, so it does not count and numbers that 
come immediately after a 13 also do not count.

sum13([1, 2, 2, 1]) → 6
sum13([1, 1]) → 2
sum13([1, 2, 2, 1, 13]) → 6
*/

package main

import (
	"fmt"
)

//returns sum of slice
func sumArr(nums[]int) (result int){
	for _, val := range nums{
		result += val
	}
	return result
}

//alter values in original slice
//call helper function to get final sum
func sum13(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 13 {
			nums[i] = 0
			if i+1 < len(nums){
				nums[i+1] = 0
			}
		}
		
	}
	finalResult := sumArr(nums) //call helper function
	return finalResult
}

func main() {
	fmt.Println(sum13([]int{1,2,2,1})) //6
	fmt.Println(sum13([]int{1,1})) //2
	fmt.Println(sum13([]int{1,2,2,1,13})) //6
}

