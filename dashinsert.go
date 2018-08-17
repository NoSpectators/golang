/*
have the function DashInsert(str) insert dashes ('-') 
between each two odd numbers in str. For example: 
if str is 454793 the output should be 4547-9-3. 
Don't count zero as an odd number. 
*/

package main

import (
	"fmt"
	"strconv"
)


func DashInsert(s string) (result string) {
	for i := 0; i < len(s)-1; i++ {
		temp, err := strconv.Atoi(string(s[i]))
		temp2, err := strconv.Atoi(string(s[i+1]))
		if err == nil {
			if temp % 2 == 1 && temp2 % 2 == 1{
				result += string(s[i])
				result += "-"
			} else {
				result += string(s[i])
			}	
		} 				
	}
	result += string(s[len(s)-1])	
	return result
}

func main() {
	fmt.Println(DashInsert("454793")) //4547-9-3
	fmt.Println(DashInsert("40257")) //4025-7
}