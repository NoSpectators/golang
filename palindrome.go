package main

import (
	"fmt"
	"unicode"
)

//can a string be read the same forward and backward
//ignores non-letters and letter case
func IsPalindrome(s string) string {	

	//much faster,less memory allocations than loop below
	letters := make([]rune, 0, len(s))
	for _, char := range s{
		if unicode.IsLetter(char){
			letters = append(letters, unicode.ToLower(char))
		}
	}	
	for i := 0; i < len(letters)/2; i++ {
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

/* slower than lines 13-18, allocates more memory
letters := []string{} //get only lowercase letters
for _, char := range s {
	if unicode.IsLetter(rune(char)){
		letters = append(letters, string(unicode.ToLower(char)))
	}
}*/