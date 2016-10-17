package main

import (
	    "fmt"
	    "log"
	    "os"
        "bufio"
	    "github.com/dchest/stemmer/porter2"
       )


func check(e error) {

	if e != nil {
		log.Fatal(e)
	}
}

func main() {

    // Open file for reading from the input file
    fmt.Println("Opening file to read...")
    infile, e := os.Open("Stem_test.txt")
    check(e)

    english_stemmer := porter2.Stemmer
    
    // Open file for writing the result
    outfile, e := os.Create("Stem_test_out.txt")
    check(e)

    // Initialise the scanner Line by Line reader
    scanner := bufio.NewScanner(infile)

    // Scan the file line by line
    // get the text by .Text()
    // pass the string to stemmer object
    for scanner.Scan() {
        outfile.WriteString(english_stemmer.Stem(scanner.Text()))
        outfile.WriteString("\n")
    }
    fmt.Println("Written to file")
    
}