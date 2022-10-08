package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	filePath := "/user/local/filename.ext"

	// Open the file
	f, err := os.Open(filePath)
	if err != nil {
		log.Println("Failed to open file:", err)
		return
	}
	// Close the file before exiting the program
	defer f.Close()

	// Create a new scanner
	scanner := bufio.NewScanner(f)

	// Iterate over the file line by line
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
