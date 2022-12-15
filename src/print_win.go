package src

import (
	"fmt"
	"os"
)

func (h *HangManData) PrintWin() { // Print the win message
	fmt.Println("The word was", h.ToFind, ", congratulations")
	os.Exit(0)
}
