package src

func NewHangManData() *HangManData {
	var data *HangManData = new(HangManData)

	data.Attempts = 10
	data.MaxAttempts = 10
	data.FindWordInFile()
	data.MapMakerForToFind()
	data.ReplaceWordByUnderscore()
	data.ReavealOneLetter()

	return data
}
