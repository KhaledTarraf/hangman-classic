package src

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func (h *HangManData) FindWordInFile() { // Find a word in the file "dicti.txt"

	data, err := os.Open("dictio.txt") // Open the file

	if err != nil { // Check for errors
		fmt.Println(err) // If there is an error, print it
	}
	scanner := bufio.NewScanner(data) // Create a new scanner

	scanner.Split(bufio.ScanWords) // Split the file by words

	for scanner.Scan() { // Scan the file
		h.HangmanPositions = append(h.HangmanPositions, scanner.Text()) // Append the words to the slice
	}

	rand.Seed(time.Now().Unix())
	result := rand.Intn(len(h.HangmanPositions) - 1)

	h.ToFind = h.HangmanPositions[result]
	data.Close()
}

func (h *HangManData) ReplaceWordByUnderscore() { // Replace the word by underscores
	for range h.ToFind {
		h.Word += "_"
	}
}
