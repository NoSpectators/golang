/*Have the function ThirdGreatest(strArr) take the array of strings
 stored in strArr and return the third largest word within it. 
 So for example: if strArr is ["hello", "world", "before", "all"] 
 your output should be world because "before" is 6 letters long, 
 and "hello" and "world" are both 5, but the output should be world 
 because it appeared as the last 5 letter word in the array. 
 If strArr was ["hello", "world", "after", "all"] the output should 
 be after because the first three words are all 5 letters long, 
 so return the last one. The array will have at least three strings 
 and each string will only contain letters.
 */


package main

import (
	"fmt"
	"sort"
)

// In order to sort by a custom function in Go, we need a
// corresponding type. Here we've created a `ByLength`
// type that is just an alias for the builtin `[]string`
// type.
type ByLength []string


// We implement `sort.Interface` - `Len`, `Less`, and
// `Swap` - on our type so we can use the `sort` package's
// generic `Sort` function. `Len` and `Swap`
// will usually be similar across types and `Less` will
// hold the actual custom sorting logic. In our case we
// want to sort in order of increasing string length, so
// we use `len(s[i])` and `len(s[j])` here.
func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}


// With all of this in place, we can now implement our
// custom sort by casting the original `ex` slice to
// `ByLength`, and then use `sort.Sort` on that typed
// slice.
func main() {

	ex1 := []string{"hello", "world", "after", "all"}//after
	sort.Sort(sort.Reverse(ByLength(ex1)))
	fmt.Println(ex1[2])
	ex2 := []string{"hello","world","before","all"}//world
	sort.Sort(sort.Reverse(ByLength(ex2)))
	fmt.Println(ex2[2])
}

