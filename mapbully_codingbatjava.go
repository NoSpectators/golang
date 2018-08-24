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

		//if there is no b key
		if val, ok := mapExample["b"]; !ok {
			//fmt.Println("no b key",val)
			val = newMap["a"]
			newMap["b"] = val
			newMap["a"] = ""
		} 

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
	
	return newMap
}


func main(){
	map1 := make(map[string]string)
	map1["a"] = "candy"
	map1["b"] = "dirt"
	fmt.Println(mapBully(map1))
	
	map2 := make(map[string]string)
	map2["a"] = "candy"
	fmt.Println(mapBully(map2))

	map3 := make(map[string]string)
	map3["a"] = "candy"
	map3["b"] = "carrot"
	map3["c"] = "meh"
	fmt.Println(mapBully(map3))

	map4 := make(map[string]string)
	map4["b"] = "carrot"
	fmt.Println(mapBully(map4))

	map5 := make(map[string]string)
	map5["c"] = "meh"
	fmt.Println(mapBully(map5))

	map6 := make(map[string]string)
	map6["a"] = "sparkle"
	map6["c"] = "meh"
	fmt.Println(mapBully(map6))

	//fmt.Printf("%t",map1)
	
}
