/*
Have the function NumberSearch(str) take the str parameter,
 search for all the numbers in the string, add them together, 
 then return that final number. For example: 
 if str is "88Hello 3World!" the output should be 91. 
 You will have to differentiate between single digit numbers
 and multiple digit numbers like in the example above. 
 So "55Hello" and "5Hello 5" should return two different answers.
 Each string will contain at least one letter or symbol. 
 */


package main

import (
	"fmt"
	"strings"
	"strconv" //https://golang.org/pkg/strconv/
)


func NumberAddition(s string) (result int) {	

	strArr := strings.Split(s,"")
	tempString := ""
	for i := 0; i < len(strArr); i++ {
		char1,err := strconv.Atoi(strArr[i]) //Atoi=string to Int
		if err == nil {
			tempString += strconv.Itoa(char1)
		} else {
			tempString += ":"
		}		
	}
	strArr2 := strings.Split(tempString, ":")
	for _, val := range strArr2 {
		val, err := strconv.Atoi(val)
		if err == nil{
			result += val
		}
	}
	return result	
}

func main() {
	fmt.Println(NumberAddition("88Hello 3World!"))//91
	fmt.Println(NumberAddition("55Hello"))//55
	fmt.Println(NumberAddition("5Hello55"))//60
}

