package main

import (
	"fmt"
)

//how to print the length of an item in a map[string]interface
func main() {
	data := map[string]interface{}{
		"password":"testpassword",
		
	}
	num := data["password"].(string)
	
	fmt.Println(len(num))
}
