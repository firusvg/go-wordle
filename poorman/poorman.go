package poorman

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"go-wordle/utility"
	"go-wordle/wordlist"
)

func transformToPoorman(s string) string {
	var result = []rune{'?', '?', '?', '?', '?'}
	slice := []rune(s)

	for i := range slice {
		switch slice[i] {
		case '🟩':
			result[i] = '*'
		case '🟨':
			result[i] = '+'
		case '⬛':
			result[i] = '_'
		}
	}
	return string(result)
}

func Main() {
	wordleNo, wordleWord := wordlist.TodaysWordle()
	tryNumber := 0
	var finished bool
	var results [6]string
	var input string
	var guessStatus int
	var resultEmoji string
	var tries string

	header := "Wordle " + strconv.Itoa(wordleNo)
	fmt.Printf("\n%s\n\n", header)
	for !(finished) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">")
		input, _ = reader.ReadString('\n')
		// convert CRLF to LF for Windows
		input = strings.Replace(input, "\n", "", -1)
		if len([]rune(input)) != 5 {
			fmt.Println("Not in word list")
			continue
		}
		guessStatus, _, resultEmoji = utility.CheckWord(strings.ToUpper(input), wordleWord)
		if guessStatus < 0 {
			fmt.Println("Not in word list")
			continue
		}
		tryNumber++
		results[tryNumber-1] = resultEmoji
		switch guessStatus {
		case 0:
			finished = tryNumber == 6
		case 1:
			finished = true
		}
		fmt.Println(transformToPoorman(resultEmoji))
	}
	if guessStatus == 0 {
		fmt.Printf("\nFAIL - %s\n", strings.ToUpper(wordleWord))
		tries = "X/6\n"
	} else {
		fmt.Printf("\nWIN - %s\n", utility.WinStatuses[tryNumber-1])
		tries = fmt.Sprintf("%d/6\n", tryNumber)
	}
	fmt.Printf("\n%s %s\n", header, tries)
	for i := 0; i < tryNumber; i++ {
		fmt.Println(transformToPoorman(results[i]))
	}
}
