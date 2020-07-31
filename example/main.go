package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/pec1985/go-async"
)

func main() {
	a := async.New(10)
	words := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce vestibulum metus id interdum dapibus. Phasellus imperdiet ac tellus et porttitor"
	slice := strings.Split(words, " ")
	var newslice []string
	var mu sync.Mutex
	for _, _word := range slice {
		// IMPORTANT!
		// copy the values that need to be used inside the function to new vars
		word := _word
		a.Do(func() error {
			mu.Lock()
			newslice = append(newslice, word)
			mu.Unlock()
			return nil
		})
	}
	if err := a.Wait(); err != nil {
		panic(err)
	}
	fmt.Println(strings.Join(newslice, ", "))
}
