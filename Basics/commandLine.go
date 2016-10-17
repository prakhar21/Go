package main

import (
	   "fmt"
	   "os"
       )


func main(){
	
    arguments := os.Args[1:]
    
    fmt.Println("argument[0]",arguments[0])
    fmt.Println("argument[1]",arguments[1])
    fmt.Println("argument[2]",arguments[2])

}
