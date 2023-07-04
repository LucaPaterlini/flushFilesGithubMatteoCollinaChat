package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// Write a file
	filePath := "foo.json"

	// Data to write to the file
	data := []byte("{\n  \"greating\": \"Hello\"\n}")

	// Write the data to the file
	err := os.WriteFile(filePath, data, 0644)

	// Communicate the file is ready for the read
	resp, err := http.Get("http://localhost:8888?filename=foo.json")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
}
