package main

//import "os"
import "time"
import "fmt"

/*
 Main function for the program. It will run through a given text file and scan all of the words, lines, and characters and return statistics about them.
*/
func main(){
  filename:=os.Args[1]
  var c chan string = make(chan string)
  //Create the concurrency. When go is called, the method is run and the next line starts executing. This will start
  //the scanning of words, characters, and lines running at the same time
  go scanWords(c)
  go scanCharacters(c)
  go scanLines(c)
  go printer(c)
  for{
    
  }
}

func printer(c chan string) {
  for {
    msg := <- c
    fmt.Println(msg)
    time.Sleep(1100)
  }
}

/*
 This function will do the scanning of words and print the statistics associated with it.
*/ 
func scanWords(c chan string){
  for x:=0; x<=100; x++{
    c <- "words"
  }
}

/*
 This function will do the scanning of characters and print the statistics associated with it.
*/ 
func scanCharacters(c chan string){
  for x:=0; x<=100; x++{
    c <- "characters"
  }
}

/*
 This function will do the scanning of lines and print the statistics associated with it.
*/ 
func scanLines(c chan string){
  for x:=0; x<=100; x++{
    c <- "lines"
  }
}