package main

import (
	"fmt"
	"net/http"
	"os"

	examplegodependency "github.com/markcanfly/example-go-dependency"
)

func main() {
	fmt.Println("Hello, World!")

	fmt.Println(examplegodependency.SpongeBobCase("I am a secure dependency"))

	// CAPABILITY_NETWORK
	_, err := http.Get("http://example.org")
	if err != nil {
		fmt.Printf("Error making HTTP request: %v\n", err)
	}

	// CAPABILITY_READ_SYSTEM_STATE
	cwd, _ := os.Getwd()
	fmt.Println("Current working directory:", cwd)
}
