package src

import (
	"fmt"
	"os"
)

func (h *HangManData) PrintLose() { // Print the lose message
	fmt.Println("The word was", h.ToFind, ", I'm sorry but you lose")
	os.Exit(0)
}
