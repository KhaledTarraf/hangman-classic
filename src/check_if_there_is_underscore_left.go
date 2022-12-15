package src

func (h *HangManData) CheckIfThereIsUnderscorLeft() bool { // Check if there is an underscore in the word
	var underscoreLeft bool
	for _, v := range h.Word {
		if string(v) == "_" {
			underscoreLeft = true
		}
	}
	return underscoreLeft
}
