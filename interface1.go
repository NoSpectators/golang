package main

import (
    "fmt"
)


//struct 
type User struct {
    FirstName, LastName string
}

//ensures type is "User", returns string or error 
func (u *User) Name() string {
    return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

//interface is a set of  methods; 
//a value of interface type can hold any value that implements those methods
type Namer interface {
    Name() string
}

func Greet(n Namer) string {
    return fmt.Sprintf("Dear %s", n.Name())
}

func main() {
    u := &User{"Matt", "Aimonetti"} //passes strings as a pointer to user struct
    fmt.Println(Greet(u))
}

