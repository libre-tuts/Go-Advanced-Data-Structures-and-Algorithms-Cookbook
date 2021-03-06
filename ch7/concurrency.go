//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt,sort, errors, math/rand and strings packages

import (
	"fmt"
	"time"
)

// main method
func main() {
	var theLibrary []string
	theLibrary = []string{"cup", "book", "book", "cup", "book"}
	var bookChannel chan string
	bookChannel = make(chan string)
	var readerChannel chan string
	readerChannel = make(chan string)

	go func(library []string) {
		var item string
		for _, item = range library {
			if item == "book" {
				bookChannel <- item
			}
		}
	}(theLibrary)

	go func() {
		var i int
		for i = 0; i < 3; i++ {
			var foundBook string
			foundBook = <-bookChannel
			fmt.Println("From Finder: ", foundBook)
			readerChannel <- "read Book"
		}
	}()

	go func() {
		var i int
		for i = 0; i < 3; i++ {
			var readBook string
			readBook = <-readerChannel
			fmt.Println("From Reader: ", readBook)
			fmt.Println("From Librarian: Book is Returned")
		}
	}()
	<-time.After(time.Second * 5)

}
