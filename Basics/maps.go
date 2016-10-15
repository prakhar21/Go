package main

import (
           "fmt"
           "reflect"
       )

func main(){
    
    // Maps are Key, Value pairs just like Dictionaries in Python
    myMap := make(map[string] int)
    
    myMap["test"] = 20

    fmt.Println("myMap is of type :",reflect.TypeOf(myMap))
    fmt.Println(myMap["test"])
}
