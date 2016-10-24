package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
       )
   

func check(e error) {

    if e != nil {
        panic(e)
    }

}

func main() {

    urlEndpoint := "https://en.wikipedia.org/wiki/Precision_and_recall"
    
    // return the response object and error
    response, err := http.Get(urlEndpoint)
    check(err)

    defer response.Body.Close()
 
    // read the contents of the body for the response
    bytes, err := ioutil.ReadAll(response.Body)
    check(err)

    // convert the ioutil byte array to string
    content := string(bytes)

    fmt.Println(content)
}
