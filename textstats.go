package main

import "os"
import "fmt"

/*
 Main function for the program. It will run through a given text file and scan all of the words, lines, and characters and return statistics about them.
*/
func main(){
  filename:=os.Args[1]
  
  //Create the concurrency. When go is called, the method is run and the next line starts executing. This will start
  //the scanning of words, characters, and lines running at the same time
  go scanWords()
  go scanCharacters()
  go scanLines()
}

/*
 This function will do the scanning of words and print the statistics associated with it.
*/ 
func scanWords(file String){
  
}

/*
 This function will do the scanning of characters and print the statistics associated with it.
*/ 
func scanCharacters(file String){
  
}

/*
 This function will do the scanning of lines and print the statistics associated with it.
*/ 
func scanLines(file String){
  
}