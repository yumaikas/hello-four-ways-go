package main

// This is a gratuitous use of goroutines, don't do this in production
import (
	"fmt"
	"time"
)

// 4.1 OMIT

func main() {

	stringChan := make(chan string) // Insert anime joke
	go func() {
		stringChan <- "Hello, "
		time.Sleep(3 * time.Second)
		stringChan <- "World!\n"
		close(stringChan)
	}()
	printBufferedChars(stringChan)
}

func printBufferedChars(stringsChannel chan string) {
	for str := range stringsChannel {
		fmt.Print(str)
	}
}

// 4.1end OMIT
