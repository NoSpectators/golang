/*Return an array that contains exactly the same numbers as the given array, 
but rearranged so that every 3 is immediately followed by a 4. Do not move the 3's, 
but every other number may move. The array contains the same number of 3's and 4's, 
every 3 has a number after it that is not a 3, and a 3 appears in the array before any 4.

fix34([1, 3, 1, 4]) → [1, 3, 4, 1]
fix34([1, 3, 1, 4, 4, 3, 1]) → [1, 3, 4, 1, 1, 3, 4]
fix34([3, 2, 2, 4]) → [3, 4, 2, 2]


*/

package main 

import (
	"fmt"
)

func fix34(arr []int) ([]int){
	
	for i:= 0; i < len(arr); i++ {
		if arr[i] == 3 {
			temp := arr[i+1] //saves value at i+1
			arr[i+1] = 4
			for j := i+2; j < len(arr); j++{
				if arr[j] == 4{
					arr[j] = temp //replaces 4 with temp
				}
			}
		}
	}
	return arr
}

func main() {

	fmt.Println(fix34([]int{1,3,1,4}))	//[1 3 4 1]
	fmt.Println(fix34([]int{1,3,1,4,4,3,1})) //[1, 3, 4, 1, 1, 3, 4]

}