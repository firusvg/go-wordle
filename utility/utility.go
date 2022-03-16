package utility

import (
	"strings"

	"go-wordle/wordlist"
)

var WinStatuses = []string{"Genius", "Magnificent", "Impressive", "Splendid", "Great", "Phew"}

func IndexOfGen[T rune | string](slice []T, item T) int {
	for i := range slice {
		if slice[i] == item {
			return i
		}
	}
	return -1
}
func CheckWord(_wordToCheck string, _wordGoal string) (result int, resultString string, resultEmoji string) {
	// result: -1 - invalid word
	//          0 - not
	//          1 - new try
	var letters [5]string
	resultString = ""
	resultEmoji = ""
	wordToCheck := []rune(_wordToCheck)
	wordGoal := []rune(_wordGoal)

	//Is word valid?
	if (IndexOfGen(wordlist.Solutions, strings.ToLower(_wordToCheck)) < 0) && (IndexOfGen(wordlist.AcceptedWords, strings.ToLower(_wordToCheck)) < 0) {
		resultString = "<wr>" + _wordToCheck + "<w0>"
		resultEmoji = "🟥🟥🟥🟥🟥"
		return -1, resultString, resultEmoji
	}

	//Is word guessed?
	if _wordToCheck == _wordGoal {
		resultString = "<bg>" + _wordToCheck + "<w0>"
		resultEmoji = "🟩🟩🟩🟩🟩"
		return 1, resultString, resultEmoji
	}

	//Check green letters
	for i := 0; i < len(wordToCheck); i++ {
		if wordToCheck[i] == wordGoal[i] {
			letters[i] = "<bg>" + string(wordToCheck[i]) + "<w0>"
			wordToCheck[i] = '🟩'
			wordGoal[i] = '🟩'
		}
	}

	//Check yellow and wrong letters
	for i := 0; i < len(wordToCheck); i++ {
		if wordToCheck[i] != '🟩' {
			j := IndexOfGen(wordGoal, wordToCheck[i])
			if j >= 0 {
				letters[i] = "<by>" + string(wordToCheck[i]) + "<w0>"
				wordToCheck[i] = '🟨'
				wordGoal[j] = '🟩'
			} else {
				letters[i] = string(wordToCheck[i])
				wordToCheck[i] = '⬛'
			}
		}
	}

	for i := 0; i < 5; i++ {
		resultString += letters[i]
		resultEmoji += string(wordToCheck[i])
	}
	return 0, resultString, resultEmoji
}
