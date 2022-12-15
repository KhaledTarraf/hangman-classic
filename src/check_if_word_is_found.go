package src

import "strings"

func (h *HangManData) CheckIfWordIsFound() { // Check if the word is found
	if strings.Compare(h.Letter, h.ToFind) == 0 {
		h.PrintWin()
	} else {
		h.Attempts -= 2
	}
}
