package main

import "fmt"
//import "reflect"

func main(){

    // Simplified form of below mentioned looping style
    fmt.Println("First Style")
    i := 0
    for i<=10{
        fmt.Println("Value of i is:",i)
	i += 1
    }
 
    // Loop like most of the languages 
    fmt.Println("Second Style")
    for i:=0; i<5; i++ {
	fmt.Println("Value of i is:",i)        

    }

  
}
