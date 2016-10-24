package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	// Create a file and write down some text
	file, err := os.Create("Sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("This is a sample text.")

	file.Close()

	// Open file for reading
	stream, err := ioutil.ReadFile("Sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	readLine := string(stream)
	fmt.Println(readLine)
}
