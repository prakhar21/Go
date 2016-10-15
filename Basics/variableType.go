package main

// Importing Packages
import (
           "fmt"
           "reflect"
       )

func main(){
    // variable decl.
    var (
	    x int
	    y float32
	    z bool
            myName string
	)

    // Print out the type of the variable
    fmt.Println("Type of x is",reflect.TypeOf(x))
    fmt.Println("Type of y is",reflect.TypeOf(y))
    fmt.Println("Type of z is",reflect.TypeOf(z))

    // Variable assignment
    // Print the value of the variable
    x,y = 10,12.233
    fmt.Println(x,y)   
    fmt.Println("Default value of z is :",z)
 
    // Constant value 
    const pi float64 = 3.14562
    fmt.Println("Value of Pi is :", pi)
    // Trying to change the value (uncomment below line to test)
    // pi = 12.2212
    // Should throw error
    
    // len(<string>) can be used to find length of the string
    myName = "test"
    fmt.Println("My name is",myName,"and it has :",len(myName),"char.")

}
