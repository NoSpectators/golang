package main

import (
	"fmt"
	"unicode"
)

//can a string be read the same forward and backward
//ignores non-letters and letter case
func IsPalindrome(s string) string {	

	letters := []string{} //get only lowercase letters

	for _, char := range s {
		if unicode.IsLetter(rune(char)){
			letters = append(letters, string(unicode.ToLower(char)))
		}
	}
	for i := 0; i < len(letters); i++ {
		if letters[i] != letters[len(letters)-i-1] {
			return "false"
		}
	}
	return "true"	
}



func main() {

	fmt.Println(IsPalindrome("racecaR"))//true
	fmt.Println(IsPalindrome("foxof")) //true
	fmt.Println(IsPalindrome("racecer")) //false

}

