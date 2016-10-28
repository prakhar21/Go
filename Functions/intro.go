package main


import (
	"fmt"
)

func main() {
    
    sum := addNumbers(1,2,3,4,5)

    fmt.Println("Sum is :",sum)

}

//... passses multiple value of same datatype as the parameter to the function
// Functions starting with small letter are private to some package
// Starting with capital letter exposes it as public functions.
func addNumbers(values ...int) int {
    sum := 0

    for i := range values {
        sum = sum + values[i]
    }

    return sum
}
