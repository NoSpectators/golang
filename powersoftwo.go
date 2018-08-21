/*
Have the function PowersofTwo(num) take the num parameter 
being passed which will be an integer and return the string true
if it's a power of two. If it's not return the string false.
For example if the input is 16 then your program should return 
the string true but if the input is 22 then the output should 
be the string false. 
*/

package main

import (
	"fmt"
)

/*
use properties of `bitwise and` to determine answer
see https://www.exploringbinary.com/ten-ways-to-check-if-an-integer-is-a-power-of-two-in-c/
*/
func PowersOfTwo(num int) string {
	fmt.Printf("%08b\n",num)
	fmt.Printf("%08b\n",num-1)
	if num & (num-1) == 0 {
		return "true"
	}
	return "false"
}

func main() {
	fmt.Println(PowersOfTwo(16))//true
	fmt.Println(PowersOfTwo(22))//false
}

