package main 

import (
			"fmt"
			"github.com/levigross/grequests"
			"log"
			"os"
			"bufio"
	   )


func check(e error) {

	if e != nil {
		log.Fatal(e)
	}
}

func main() {

    // Open file to read url from there
	fmt.Println("Opening File...")
	infile, err := os.Open("urls.txt")
	check(err)

	// Initialise the scanner Line by Line reader
    scanner := bufio.NewScanner(infile)

    for scanner.Scan() {
    	resp, err := grequests.Get(scanner.Text(), nil)
    	check(err)

    	if resp.Ok != true {
    		log.Fatal('Response is not Ok!')
    	}

    	fmt.Println(resp.String())
    }
}

