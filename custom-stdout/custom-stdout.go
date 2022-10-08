package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func print() {
	fmt.Println("Printed using the new pipe")
	fmt.Println("This is another new line printed using my custom stdout")
}

func main() {
	// Create a backup of original stdout
	readerBackup := os.Stdout

	// Create a new stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Now everything you print gets sent to your new stdout
	print()

	// Create a channel to print the data in the std out
	outC := make(chan string)

	// Copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	// Close the custom stdout
	w.Close()

	// Restore the original stdout
	os.Stdout = readerBackup
	out := <-outC

	// Reading data from our custom stdout
	fmt.Println("Back to the original stdout state")
	fmt.Println("Printing data from the channel...")
	fmt.Print(out)
}

