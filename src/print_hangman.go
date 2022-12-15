package src

import (
	"fmt"
	"os"
)

func (h *HangManData) hangmansFromBytes(b []byte) []string {
	var arr []string = make([]string, 0)
	var tmp string
	arr = append(arr, "")
	for index, _byte := range b {
		if index != 0 && index%71 == 0 {
			arr = append(arr, tmp)
			tmp = ""
		}
		tmp += string(_byte)
	}
	arr = append(arr, tmp)
	return arr
}

func (h *HangManData) ReadHangmansFile() []string {
	var filename string = "hangman.txt"
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error parsing file %s\n", filename)
		fmt.Println(err)
		os.Exit(0)
	}
	return h.hangmansFromBytes(content)
}

func (h *HangManData) PrintHangman() {
	var hangmans []string = h.ReadHangmansFile()
	fmt.Println(hangmans[h.MaxAttempts-h.Attempts])
}
