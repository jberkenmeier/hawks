package main

//import "os"
import (
	"bufio"
	"fmt"
	"os"
	//"time"
	"strings"

)

var p = fmt.Print

/*
 Main function for the program. It will run through a given text file and scan all of the words, lines, and characters and return statistics about them.
*/
func main() {
	//filename := os.Args[1]
	var c chan int = make(chan int)
	var a chan [24]int = make(chan [24]int)
	var b chan [26]int = make(chan [26]int)
	var s chan string = make(chan string)
	//Create the concurrency. When go is called, the method is run and the next line starts executing. This will start
	//the scanning of words, characters, and lines running at the same time

	go scanWords(c,s)
	go scanCharacters(c,s)
	go scanLines(c,s)
	go scanWordLengths(a,s)
	go scanAvgWordLength(c,s)
	go scanLetterCount(b,s)
	go printer(c,a,b,s)
	for {

	}
}

func printer(c chan int, a chan [24]int, b chan [26]int, s chan string) {
	for {
		select {
		case msg := <-c:
		    p(<-s)
		    fmt.Println(msg)
		case arr := <-b:
			p(<-s)
			fmt.Printf("# of a's: %d\n", arr[0])
			fmt.Printf("# of b's: %d\n", arr[1])
			fmt.Printf("# of c's: %d\n", arr[2])
			fmt.Printf("# of d's: %d\n", arr[3])
			fmt.Printf("# of e's: %d\n", arr[4])
			fmt.Printf("# of f's: %d\n", arr[5])
			fmt.Printf("# of g's: %d\n", arr[6])
			fmt.Printf("# of h's: %d\n", arr[7])
			fmt.Printf("# of i's: %d\n", arr[8])
			fmt.Printf("# of j's: %d\n", arr[9])
			fmt.Printf("# of k's: %d\n", arr[10])
			fmt.Printf("# of l's: %d\n", arr[11])
			fmt.Printf("# of m's: %d\n", arr[12])
			fmt.Printf("# of n's: %d\n", arr[13])
			fmt.Printf("# of o's: %d\n", arr[14])
			fmt.Printf("# of p's: %d\n", arr[15])
			fmt.Printf("# of q's: %d\n", arr[16])
			fmt.Printf("# of r's: %d\n", arr[17])
			fmt.Printf("# of s's: %d\n", arr[18])
			fmt.Printf("# of t's: %d\n", arr[19])
			fmt.Printf("# of u's: %d\n", arr[20])
			fmt.Printf("# of v's: %d\n", arr[21])
			fmt.Printf("# of w's: %d\n", arr[22])
			fmt.Printf("# of x's: %d\n", arr[23])
			fmt.Printf("# of y's: %d\n", arr[24])
			fmt.Printf("# of z's: %d\n", arr[25])
		case array := <-a:
			p(<-s)
			fmt.Printf("Words with length of 1: %d\n", array[0])
			fmt.Printf("Words with length of 2: %d\n", array[1])
			fmt.Printf("Words with length of 3: %d\n", array[2])
			fmt.Printf("Words with length of 4: %d\n", array[3])
			fmt.Printf("Words with length of 5: %d\n", array[4])
			fmt.Printf("Words with length of 6: %d\n", array[5])
			fmt.Printf("Words with length of 7: %d\n", array[6])
			fmt.Printf("Words with length of 8: %d\n", array[7])
			fmt.Printf("Words with length of 9: %d\n", array[8])
			fmt.Printf("Words with length of 10: %d\n", array[9])
			fmt.Printf("Words with length of 11: %d\n", array[10])
			fmt.Printf("Words with length of 12: %d\n", array[11])
			fmt.Printf("Words with length of 13: %d\n", array[12])
			fmt.Printf("Words with length of 14: %d\n", array[13])
			fmt.Printf("Words with length of 15: %d\n", array[14])
			fmt.Printf("Words with length of 16: %d\n", array[15])
			fmt.Printf("Words with length of 17: %d\n", array[16])
			fmt.Printf("Words with length of 18: %d\n", array[17])
			fmt.Printf("Words with length of 19: %d\n", array[18])
			fmt.Printf("Words with length of 20: %d\n", array[19])
			fmt.Printf("Words with length of 21: %d\n", array[20])
			fmt.Printf("Words with length of 22: %d\n", array[21])
			fmt.Printf("Words with length of 23: %d\n", array[22])
			fmt.Printf("Words with length of 24: %d\n", array[23])
		default:
			
		}
	}
}

/*
 This function will do the scanning of words and print the statistics associated with it.
*/
func scanWords(c chan int, s chan string) {
	var wordcount = 0
	file, _ := os.Open("alice.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		wordcount++
	}
	c <- wordcount
	s <- "WordCount: \n"

}

/*
 This function will do the scanning of characters and print the statistics associated with it.
*/
func scanCharacters(c chan int, s chan string) {
	var charcount = 0
	file, _ := os.Open("alice.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		charcount++
	}
	c <- charcount
	s <-"CharacterCount: \n"
}

/*
 This function will do the scanning of lines and print the statistics associated with it.
*/
func scanLines(c chan int, s chan string) {
	var linecount = 0
	file, _ := os.Open("alice.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linecount++
	}
	c <- linecount
	s <- "LineCount: \n"
}

func scanWordLengths(c chan [24]int, s chan string){
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
	c <- wordLength
	s <- "WordLengths: \n"
}

func scanAvgWordLength(c chan int, s chan string){
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
	c <- int(count)
	s <- "Average Word Lengths: \n"
}


func scanLetterCount(c chan [26]int, s chan string){
	var letters [26]int
	//var chars byte 
	file, _ := os.Open("alice.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		lowerstring := strings.ToLower(scanner.Text())
		bit := []byte(lowerstring)
		if ('a' <= bit[0] && bit[0]<= 'z'){
			letters[bit[0] - 'a']++
		}
	}
	c <- letters
	s <- "Letter Count: \n"
}