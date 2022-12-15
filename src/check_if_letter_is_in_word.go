package src

import "fmt"

func (h *HangManData) CheckIfLetterIsInWord() { // Check if the letter is in the word
	for _, v := range h.ToFind {
		if h.Letter == string(v) {
			fmt.Println("The letter", h.Letter, "is in the word")
			h.DiscoveredLetters[h.Letter] = true
			h.ReplaceUnderscoreByLetter()
			h.AskForALetter()
		}
	}
	h.Attempts--
}
