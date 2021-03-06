//https://www.devdungeon.com/content/web-scraping-go
package main

import (
    "io"
    "log"
    "net/http"
    "os"
)

func main() {
    // Make HTTP GET request
    response, err := http.Get("https://www.devdungeon.com/")
    if err != nil {
        log.Fatal(err)
    }
    /*defer is often used with paired operations like 
    open/close, connect/disconnect, lock/unlock
    ensures resources are released in all cases
    the right place for a defer statement is immediately 
    after the resource has been acquired
    */
    defer response.Body.Close()

    // Copy data from the response to standard output
    n, err := io.Copy(os.Stdout, response.Body)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Number of bytes copied to STDOUT:", n)
}
