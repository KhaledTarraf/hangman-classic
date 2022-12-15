package src

func (h *HangManData) ReplaceUnderscoreByLetter() { // Replace the underscore by the letter
	for i, v := range h.ToFind {
		for k := range h.Word {
			if k == i && string(v) == h.Letter {
				h.Word = h.Word[:k] + h.Letter + h.Word[k+1:]
			}
		}
	}
}
