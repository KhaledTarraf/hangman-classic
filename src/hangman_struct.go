package src

type HangManData struct {
	Word              string          // Word composed of '_', ex: H_ll_
	ToFind            string          // Final word chosen by the program at the beginning. It is the word to find
	Attempts          int             // Number of attempts left
	HangmanPositions  []string        // It can be the array where the positions parsed in "dictio.txt" are stored
	Letter            string          // Letter chosen by the user
	DiscoveredLetters map[string]bool // Letters discovered by the user
	MaxAttempts       int             // Max attempts allowed
}
