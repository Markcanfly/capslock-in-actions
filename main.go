package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Making a network call to example.org...")

	// Make a network request to example.org
	_, err := http.Get("http://example.org")
	if err != nil {
		fmt.Printf("Error making HTTP request: %v\n", err)
		os.Exit(1)
	}

}
