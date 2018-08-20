package word

import (
	"testing"
	"unicode"
)

//at command line, run go test -v palindrome_test.go

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

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"", "true"},
		{"a", "true"},
		{"aa", "true"},
		{"ab", "false"},
		{"kayak", "true"},
		{"detartrated", "true"},
		{"A man, a plan, a canal: Panama", "true"},
		{"Evil I did dwell; lewd did I live.", "true"},
		{"Able was I ere I saw Elba", "true"},
		{"été", "true"},
		{"Et se resservir, ivresse reste.", "true"},
		{"palindrome", "false"}, // non-palindrome
		{"desserts", "false"},   // semi-palindrome
	}
	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}
