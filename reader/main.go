package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	filePath := "foo.json"
	// Read the contents of the file
	content, _ := os.ReadFile(filePath)
	expectedContent := []byte("{\n  \"greating\": \"Hello\"\n}")
	if !bytes.Equal(content, expectedContent) {
		panic("the file was not available when asked for it")
	}
	fmt.Println("Read successfully")
	fmt.Println(string(content))
}
