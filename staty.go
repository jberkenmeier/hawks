package main

//import "os"
import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var p = fmt.Print

/*
 Main function for the program. It will run through a given text file and scan all of the words, lines, and characters and return statistics about them.
*/
func main() {
	//filename := os.Args[1]
	var c chan int = make(chan int)
	//Create the concurrency. When go is called, the method is run and the next line starts executing. This will start
	//the scanning of words, characters, and lines running at the same time

	go scanWords(c)
	go scanCharacters(c)
	go scanLines(c)
	go printer(c)
	for {

	}
}

func printer(c chan int) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(1)
	}
}

/*
 This function will do the scanning of words and print the statistics associated with it.
*/
func scanWords(c chan int) {
	var wordcount = 0
	file, _ := os.Open("alice.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		wordcount++
	}
	p("WordCount: ")
	c <- wordcount

}

/*
 This function will do the scanning of characters and print the statistics associated with it.
*/
func scanCharacters(c chan int) {
	var charcount = 0
	file, _ := os.Open("alice.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		charcount++
	}
	p("CharacterCount: ")
	c <- charcount
}

/*
 This function will do the scanning of lines and print the statistics associated with it.
*/
func scanLines(c chan int) {
	var linecount = 0
	file, _ := os.Open("alice.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linecount++
	}
	p("LineCount: ")
	c <- linecount
}

