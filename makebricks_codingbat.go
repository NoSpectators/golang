/*
We want to make a row of bricks that is goal inches long.
We have a number of small bricks (1 inch each) and big bricks (5 inches each). 
Return True if it is possible to make the goal by choosing from the given bricks. 
This is a little harder than it looks and can be done without any loops. 

make_bricks(3, 1, 8) → True
make_bricks(3, 1, 9) → False
make_bricks(3, 2, 10) → True
*/

package main

import (
	"fmt"
)

func makeBricks(nums []int) bool{

	small := nums[0]
	big := nums[1]
	goal := nums[2]

	//step 1 make sure you have at least enough bricks
	if small + big*5 >= goal {

		//step 2 ensure proportions are right of big to small
		if (goal % 5) - small <= 0 {
			return true
		}
	}
	return false 
}

func main() {
	fmt.Println(makeBricks([]int{3,1,8})) //true
	fmt.Println(makeBricks([]int{3,1,9})) //false
	fmt.Println(makeBricks([]int{43,1,46})) //true
}

