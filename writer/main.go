package main

import (
	"os"
)

func main() {
	// Write a file
	filePath := "foo.json"

	// Data to write to the file
	data := []byte("{\n  \"greating\": \"Hello\"\n}")

	// Write the data to the file
	_ = os.WriteFile(filePath, data, 0644)
}
