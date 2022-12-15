package src

import "fmt"

func (h *HangManData) PrintAttempts() { // Print the number of attempts left
	fmt.Println("You have", h.Attempts, "attempts left")
}
