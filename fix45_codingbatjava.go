/*

(This is a slightly harder version of the fix34 problem.) 
Return an array that contains exactly the same numbers as the given array, 
but rearranged so that every 4 is immediately followed by a 5. 
Do not move the 4's, but every other number may move. 
The array contains the same number of 4's and 5's, 
and every 4 has a number after it that is not a 4. 
In this version, 5's may appear anywhere in the original array.

fix45([5, 4, 9, 4, 9, 5]) → [9, 4, 5, 4, 5, 9]
fix45([1, 4, 1, 5]) → [1, 4, 5, 1]
fix45([1, 4, 1, 5, 5, 4, 1]) → [1, 4, 5, 1, 1, 4, 5]
*/

package main 

import (
	"fmt"
)

func fix45(arr []int) ([]int){

	index := 0

	for i:= 0; i < len(arr)-2; i++ {

		if arr[i] == 4 && arr[i+1] != 5{ //check for 4's with no 5's after

			//fmt.Println("found 4 with no 5 after at index", i )

			for j := index; j < len(arr); j++ { 
				if arr[j] == 5 && j == 0{  //checks arr[0] for 5
					//fmt.Println("found 5 at index 0")
					arr[0] = arr[i+1] 
					//fmt.Println("arr[i+1] = ", arr[i+1])
					arr[i+1] = 5 //this swaps the non-5 with 5 after the 4
					index++
				}
				if arr[j] == 5 && arr[j-1] != 4{ //checks all other cases
					//fmt.Println("found 5 at non-zero-index", j)
					arr[j] = arr[i+1]
					arr[i+1] = 5 //swaps non-5 with 5 
					index = j 
					break
				}
			}
		}
	}
	return arr
}

func main() {

	fmt.Println(fix45([]int{5,4,9,4,9,5}))	//[9 4 5 4 5 9]
	fmt.Println(fix45([]int{1,4,1,5})) //[1 4 5 1]

}