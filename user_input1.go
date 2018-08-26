package main 

import (

    "bufio"
    "fmt"
    "os"
)

func  main() {


    
    scanner := bufio.NewReader(os.Stdin)
    fmt.Print("Enter text: ")
    text, _ := scanner.ReadString('\n')
    fmt.Println("you typed", text)
    

    //multiple lines of input
    /*
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
    */

}


