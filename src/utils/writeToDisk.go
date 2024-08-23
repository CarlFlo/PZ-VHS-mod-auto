package utils

import (
	"fmt"
	"os"

	"github.com/CarlFlo/malm"
)

// WriteStringsFromChannel listens to a channel and writes each string to a file
func WriteStringsFromChannel(ch <-chan string, filename string, done chan<- bool) {
	// Create or open the file for writing
	file, err := os.Create(filename)
	if err != nil {
		malm.Fatal("Error creating file: '%v'", err)
		done <- false
		return
	}
	defer file.Close()

	// Listen to the channel and write each string to the file
	for str := range ch {
		_, err := fmt.Fprintln(file, str)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			done <- false
			return
		}
	}

	// Indicate completion
	done <- true
}
