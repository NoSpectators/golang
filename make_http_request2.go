// make_http_request_with_timeout.go
//taken from https://www.devdungeon.com/content/web-scraping-go#make_http_get_request
package main

import (
    "io"
    "log"
    "net/http"
    "os"
    "time"
)

func main() {
    // Create HTTP client with timeout
    //this is because there is no default timeout!!

    /*see https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779 for 
    why not to use default http client in go
    */
    client := &http.Client{
        Timeout: 10 * time.Second,
    }

    // Make request
    response, err := client.Get("https://www.devdungeon.com/")
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

    // Copy data from the response to standard output
    n, err := io.Copy(os.Stdout, response.Body)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Number of bytes copied to STDOUT:", n)
}