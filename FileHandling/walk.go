package main

import (
	"fmt"
	"os"
	"path/filepath"
       )


func check(e error) {

    if e != nil {
        panic(e)
    }
}

func main() {
      
    // Current file path (absolute path)
    root, _ := filepath.Abs(".")

    // Walk through all the files in the current directory
    err := filepath.Walk(root, processPath)
    check(err)
}

func processPath(path string, info os.FileInfo, err error) error {
    
    check(err)

    if info.IsDir() {
        fmt.Println("Directories:",path)
    } else {
        fmt.Println("Files:",path)
    }
    
    return nil  
}
