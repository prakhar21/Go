package main

//stringutil.go is defined under src/stringutil/ path
import (
	"fmt"
	"stringutil"
)


// ConcatString is a public function defined in the stringutil.go
func main() {
    s := stringutil.ConcatString("Prakhar","Mishra") 
    fmt.Println(s)
}
