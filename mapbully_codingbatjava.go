/*Modify and return the given map as follows: 
if the key "a" has a value, set the key "b" to have that value, 
and set the key "a" to have the value "". Basically "b" is a bully,
 taking the value and replacing it with the empty string.
 */

package main

import (
	"fmt"
)


func mapBully (mapExample map[string]string) (map[string]string){

	//create new target map
	newMap := make(map[string]string)	
	for k, v := range mapExample{
		newMap[k] = v
	}
	//check of map is length 1
	if len(newMap) == 1 {

		//if length 1 and has "a" key, assign b key
		if val, ok := newMap["a"]; ok {
			//fmt.Println("val",val)
			newMap["b"] = val
			newMap["a"] = ""
		} 

	} else { //order matters here

		if val, ok := mapExample["a"]; !ok { //if there is no a key but is b
			if val2, ok := mapExample["b"]; ok {
				val = ""
				newMap["a"] = val
				val2 = mapExample["b"]
				newMap["b"] = val2
			}
			
		} else if val, ok := mapExample["b"]; !ok {  //if there is no b key but is a
			if val2, ok := mapExample["a"]; ok {
				val2 = ""
				val = mapExample["a"]
				newMap["b"] = val
				newMap["a"] = val2
			}			
		} else { // both a and b keys exist 
			//replace map[b] with map[a]'s value
			if val, ok := mapExample["b"]; ok {
				val = newMap["a"]
			 	newMap["b"] = val			 		
		 	} 
		 	//replace map[a] with ""
		 	if val, ok := mapExample["a"]; ok {  
				val = ""
				newMap["a"] = val
		 	} 
		}
	}	
	return newMap
}


func main(){
	map1 := make(map[string]string)
	map1["a"] = "candy"
	map1["b"] = "dirt"
	fmt.Println("original =", map1, "solution =", mapBully(map1))
	
	map2 := make(map[string]string)
	map2["a"] = "candy"
	fmt.Println("original =", map2, "solution =", mapBully(map2))

	map3 := make(map[string]string)
	map3["a"] = "candy"
	map3["b"] = "carrot"
	map3["c"] = "meh"
	fmt.Println("original =", map3, "solution =", mapBully(map3))

	map4 := make(map[string]string)
	map4["b"] = "carrot"
	fmt.Println("original =", map4, "solution =", mapBully(map4))

	map5 := make(map[string]string)
	map5["c"] = "meh"
	fmt.Println("original =",map5, "solution =",mapBully(map5))

	map6 := make(map[string]string)
	map6["a"] = "sparkle"
	map6["c"] = "meh"
	fmt.Println("original =",map6, "solution =",mapBully(map6))

	map7 := make(map[string]string)
	map7["b"] = "yo"
	map7["c"] = "diddy"
	map7["e"] = "pow"
	map7["f"] = "not"
	fmt.Println("original =", map7, "solution =", mapBully(map7))

	//fmt.Printf("%t",map1)
	
}
