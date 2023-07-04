package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

// Initialize a WaitGroup
var wg sync.WaitGroup

func handler(w http.ResponseWriter, r *http.Request) {
	defer wg.Done()
	filePath := r.URL.Query()["filename"][0]

	// Read the contents of the file
	content, _ := os.ReadFile(filePath)
	expectedContent := []byte("{\n  \"greating\": \"Hello\"\n}")
	if !bytes.Equal(content, expectedContent) {
		panic("the file was not available when asked for it")
	}

	fmt.Fprintln(w, "All fine content:"+string(content))
}

func main() {
	// Set the number of requests to limit
	numRequests := 1

	// Add the number of requests to the WaitGroup
	wg.Add(numRequests)

	// Create an HTTP server
	server := &http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/", handler)
	fmt.Println("Running demo app. Press Ctrl+C to exit...")

	// Start the server in a goroutine
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for all requests to be processed
	wg.Wait()

	// Shutdown the server gracefully
	err := server.Shutdown(nil)
	if err != nil {
		log.Fatalf("Shutdown error: %v", err)
	}

}
