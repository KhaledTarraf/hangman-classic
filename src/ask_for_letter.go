package src

import (
	"fmt"
	"strings"
)

func (h *HangManData) AskForALetter() { // Ask for a letter
	fmt.Println("******************")
	h.PrintHangman()
	h.PrintAttempts()
	fmt.Println()
	fmt.Println(h.Word)
	fmt.Println()

	if h.CheckIfThereIsUnderscorLeft() {
		fmt.Println("Enter a letter or a word")
		fmt.Scan(&h.Letter)
		h.Letter = strings.ToUpper(strings.TrimSpace(h.Letter))

		if len(h.Letter) == 1 {
			h.CheckIfLetterIsInWord()
		} else if len(h.Letter) > 1 {
			h.CheckIfWordIsFound()
		} else {
			fmt.Println("You must enter a letter or a word")
			h.AskForALetter()
		}
	} else {
		h.PrintWin()
	}
}
