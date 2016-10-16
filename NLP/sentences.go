package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"    
    "gopkg.in/neurosnap/sentences.v1/english"
)

// Checks for the error
func check(e error) {
    
    if e != nil {
        log.Fatal(e)
    }
}

func main(){
    
    // Read text from file
    stream, err := ioutil.ReadFile("Sample.txt")    
    check(err)
    
    text := string(stream)

    tokenizer, err := english.NewSentenceTokenizer(nil)
    check(err)
  
    sentences := tokenizer.Tokenize(text)
    
    // Print the sentences
    for idx, s := range sentences {
        fmt.Println(idx, s.Text)
    }

    // Write senetences to the file
    file, err := os.Create("Sentences.txt")
    check(err)    
    
    for _, sent := range sentences {
        file.WriteString(sent.Text)
        file.WriteString("\n")	
    }
}
