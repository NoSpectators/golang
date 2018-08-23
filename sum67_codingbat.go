/*
Return the sum of the numbers in the array, except ignore sections of numbers 
starting with a 6 and extending to the next 7 
(every 6 will be followed by at least one 7). 
Return 0 for no numbers.

sum67([1, 2, 2]) → 5
sum67([1, 2, 2, 6, 99, 99, 7]) → 5
sum67([1, 1, 6, 7, 2]) → 4
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
func sum67(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 6{
			nums[i] = 0
			for j := i+1; j < len(nums); j++{
				temp := nums[j]
				nums[j] = 0
				if temp == 7 {
					i = j+1
					break;
				}
			}			
		}
	}
	finalResult := sumArr(nums) //call helper function
	return finalResult
}

func main() {
	fmt.Println(sum67([]int{1,2,2})) //5
	fmt.Println(sum67([]int{1,2,2,6,99,7})) //5
	fmt.Println(sum67([]int{1,1,6,7,2})) //4
}

