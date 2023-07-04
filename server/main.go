package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
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
	http.HandleFunc("/", handler)
	fmt.Println("Running demo app. Press Ctrl+C to exit...")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
