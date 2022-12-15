package src

func (h *HangManData) MapMakerForToFind() { // Create a map with the letters of the word to find
	h.DiscoveredLetters = make(map[string]bool)
	for _, v := range h.ToFind {
		h.DiscoveredLetters[string(v)] = false
	}
}
