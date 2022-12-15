package main

import "github.com/KhaledTarraf/hangman-classic/src"

func play() {
	hangman := src.NewHangManData()
	for hangman.Attempts > 0 {
		hangman.AskForALetter()
	}

	if hangman.Attempts == 0 {
		hangman.PrintHangman()
		hangman.PrintLose()
	}
}

func main() {
	play()
}
