package main
import (
	"fmt"
	"io"
	"os"
       )



func check(e error) {

    if e != nil {
        panic(e)
    }	
}

func main() {

    // Content to be written
    content := "This is a test example. That will be written to the file"
    
    // Returns the file object and error
    file, err := os.Create("sample.txt")
    check(err)
    defer file.Close()

    // Returns the number of chars. written to the string and error 
    lnString, err := io.WriteString(file, content)
    check(err)

    fmt.Println("Write complete!!")
    fmt.Println("Wrote", lnString,"Chars.")
}
