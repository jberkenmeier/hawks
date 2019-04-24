package main

//import "os"
import (
	"bufio"
	"fmt"
	"os"
	"time"
	"strings"
	"reflect"
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
	go scanWordLengths(c)
	go scanAvgWordLength(c)
	go scanLetterCount(c)
	go printer(c)
	for {

	}
}

func printer(c chan int) {
	for {
		msg := <-c
		rt := reflect.TypeOf(msg)
		if rt.Kind() == reflect.Array{
			if len(msg) > 24{
				fmt.Printf("# of a's: %d\n", msg[0])
				fmt.Printf("# of b's: %d\n", msg[1])
				fmt.Printf("# of c's: %d\n", msg[2])
				fmt.Printf("# of d's: %d\n", msg[3])
				fmt.Printf("# of e's: %d\n", msg[4])
				fmt.Printf("# of f's: %d\n", msg[5])
				fmt.Printf("# of g's: %d\n", msg[6])
				fmt.Printf("# of h's: %d\n", msg[7])
				fmt.Printf("# of i's: %d\n", msg[8])
				fmt.Printf("# of j's: %d\n", msg[9])
				fmt.Printf("# of k's: %d\n", msg[10])
				fmt.Printf("# of l's: %d\n", msg[11])
				fmt.Printf("# of m's: %d\n", msg[12])
				fmt.Printf("# of n's: %d\n", msg[13])
				fmt.Printf("# of o's: %d\n", msg[14])
				fmt.Printf("# of p's: %d\n", msg[15])
				fmt.Printf("# of q's: %d\n", msg[16])
				fmt.Printf("# of r's: %d\n", msg[17])
				fmt.Printf("# of s's: %d\n", msg[18])
				fmt.Printf("# of t's: %d\n", msg[19])
				fmt.Printf("# of u's: %d\n", msg[20])
				fmt.Printf("# of v's: %d\n", msg[21])
				fmt.Printf("# of w's: %d\n", msg[22])
				fmt.Printf("# of x's: %d\n", msg[23])
				fmt.Printf("# of y's: %d\n", msg[24])
				fmt.Printf("# of z's: %d\n", msg[25])
			}else{
				fmt.Printf("Words with length of 1: %d\n", msg[0])
				fmt.Printf("Words with length of 2: %d\n", msg[1])
				fmt.Printf("Words with length of 3: %d\n", msg[2])
				fmt.Printf("Words with length of 4: %d\n", msg[3])
				fmt.Printf("Words with length of 5: %d\n", msg[4])
				fmt.Printf("Words with length of 6: %d\n", msg[5])
				fmt.Printf("Words with length of 7: %d\n", msg[6])
				fmt.Printf("Words with length of 8: %d\n", msg[7])
				fmt.Printf("Words with length of 9: %d\n", msg[8])
				fmt.Printf("Words with length of 10: %d\n", msg[9])
				fmt.Printf("Words with length of 11: %d\n", msg[10])
				fmt.Printf("Words with length of 12: %d\n", msg[11])
				fmt.Printf("Words with length of 13: %d\n", msg[12])
				fmt.Printf("Words with length of 14: %d\n", msg[13])
				fmt.Printf("Words with length of 15: %d\n", msg[14])
				fmt.Printf("Words with length of 16: %d\n", msg[15])
				fmt.Printf("Words with length of 17: %d\n", msg[16])
				fmt.Printf("Words with length of 18: %d\n", msg[17])
				fmt.Printf("Words with length of 19: %d\n", msg[18])
				fmt.Printf("Words with length of 20: %d\n", msg[19])
				fmt.Printf("Words with length of 21: %d\n", msg[20])
				fmt.Printf("Words with length of 22: %d\n", msg[21])
				fmt.Printf("Words with length of 23: %d\n", msg[22])
				fmt.Printf("Words with length of 24: %d\n", msg[23])
			}
		}else{
			fmt.Println(msg)
		}
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

func scanWordLengths(c chan int){
	var wordLength [24]int
	file, _ := os.Open("alice.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		if len(scanner.Text()) < 23{
			wordLength[len(scanner.Text())]++
		}else{
			wordLength[23]++
		}
	}
	p("WordLengths: ")
	c <- wordLength
}

func scanAvgWordLength(c chan int){
	var words float64 = 0.0
	var total float64 = 0.0
	var count float64 = 0.0
	file, _ := os.Open("alice.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words += float64(len(scanner.Text()))
		total++
	}
	count = words/total
	p("Average Word Lengths: ")
	c <- count
}


func scanLetterCount(c chan int){
	var letters [26]int
	var line string
	var chars byte 
	file, _ := os.Open("alice.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = strings.ToLower(scanner.Text())
		for i, l := range line{
			i++
			letters[byte(l) - 'a']++
		}
	}
	p("Letter Count: ")
	c <- letters
}