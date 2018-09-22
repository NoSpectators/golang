/*
Have the function LongestIncreasingSequence(arr) take the array of positive integers 
stored in arr and return the length of the longest increasing subsequence (LIS). 
A LIS is a subset of the original list where the numbers are in sorted order, 
from lowest to highest, and are in increasing order. The sequence does not need to be 
contiguous or unique, and there can be several different subsequences. 
For example: if arr is [4, 3, 5, 1, 6] then a possible LIS 
is [3, 5, 6], and another is [1, 6]. For this input, your program should return 3 
because that is the length of the longest increasing subsequence. 
*/

package main 

import (
	"fmt"
)


func Max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

/*https://www.geeksforgeeks.org/longest-increasing-subsequence-dp-3/*/

func LongestIncreasingSequence (arr []int) (result int){
	
	//initialize temp array, same len as original with all 1's
	temp := make([]int, len(arr))
	for i := 0; i < len(temp); i++ {
		temp[i] = 1
	}

	//LIS algorithm
	result = 0
	for i := 1; i < len(arr); i++{
		for j := 0; j < i; j++{
			if arr[i] > arr[j] && temp[i] < temp[j]+1{
				temp[i] = temp[j]+1
			}
		}
		result = Max(result, temp[i])
		
	}
	return result
}


func main() {

	ex1 := []int{4,3,5,1,6}
	fmt.Println(LongestIncreasingSequence(ex1)) //3
	ex2 := []int{50, 3, 10, 7, 40, 80}
	fmt.Println(LongestIncreasingSequence(ex2)) //4

}