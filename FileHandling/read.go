package main

import (
	"fmt"
	"io/ioutil"
       )


func check(e error) {

    if e != nil {
        panic(e)
    }
}

func main() {

    infile := "sample.txt"
   
    // content will have the byte array for the content in the file 
    content, err := ioutil.ReadFile(infile)
    check(err)
    
    // string() function is used to convert the byte array and convert them into string array
    fmt.Println(string(content)) 
}
