/*
Have the function AdditivePersistence(num) take the num parameter being passed which will always 
be a positive integer and return its additive persistence which is the number of times you must 
add the digits in num until you reach a single digit. For example: if num is 2718 then your program 
should return 2 because 2 + 7 + 1 + 8 = 18 and 1 + 8 = 9 and you stop at 9
*/

package main

import (
	"fmt"
	"strconv"
)

func AdditivePersistence(num int) int {
	if num < 10 {
		return 0
	} else {
		temp := 0
		numstr := strconv.Itoa(num)
		//fmt.Printf("%t\n",numstr)
		for _, val := range numstr {
			val, err := strconv.Atoi(string(val)) //loop turned val to a rune
			if err == nil {
				//fmt.Printf("%t\n",val)
				temp += val
			}
		}
		return 1 + AdditivePersistence(temp)		
	}	
}

func main() {
	fmt.Println(AdditivePersistence(2718))
}

