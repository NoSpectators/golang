/*

Modify and return the given map as follows: 
if the key "a" has a value, set the key "b" to have that same value. 
In all cases remove the key "c", leaving the rest of the map unchanged.


mapShare({"a": "aaa", "b": "bbb", "c": "ccc"}) → {"a": "aaa", "b": "aaa"}
mapShare({"b": "xyz", "c": "ccc"}) → {"b": "xyz"}
mapShare({"a": "aaa", "c": "meh", "d": "hi"}) → {"a": "aaa", "b": "aaa", "d": "hi"}
 */

package main

import (
	"fmt"
)


func mapShare (mapExample map[string]string) (map[string]string){

	//create new target map
	newMap := make(map[string]string)	

	for k, v := range mapExample{
		newMap[k] = v
	}


	//if a key exists, re-assign b value to a value
	if val, ok := newMap["a"]; ok {
		newMap["b"] = val
	} else if val, ok := newMap["b"]; !ok { 
		if val2, ok := newMap["a"]; ok { //if no b key but a key exists
			val2 = newMap["a"]
			newMap["b"] = val2
		}
		val = val // if no b key don't change anything
	}

	delete(newMap, "c")
	return newMap
}


func main(){
	//{"a": "aaa", "b": "bbb", "c": "ccc"}
	map1 := make(map[string]string)
	map1["a"] = "aaa"
	map1["b"] = "bbb"
	map1["c"] = "ccc"
	fmt.Println("original =", map1, "solution =", mapShare(map1)) //map[a:aaa b:bbb]
	
	map2 := make(map[string]string)
	map2["b"] = "xyz"
	map2["c"] = "ccc"
	fmt.Println("original =", map2, "solution =", mapShare(map2)) //map[b:xyz]

	map3 := make(map[string]string)
	map3["a"] = "aaa"
	map3["c"] = "meh"
	map3["d"] = "hi"
	fmt.Println("original =", map3, "solution =", mapShare(map3)) //map[a:aaa b:aaa d:hi]

	map4 := make(map[string]string)
	map4["d"] = "456"
	fmt.Println("original =", map4, "solution =", mapShare(map4)) //map[d:456]
	
}
