package src

func (h *HangManData) ReavealOneLetter() { // Reveal one letter
	n := (len(h.ToFind) / 2) - 1
	for i, v := range h.ToFind {
		if i == n {
			h.Word = h.Word[:i] + string(v) + h.Word[i+1:]
		}
	}
}
