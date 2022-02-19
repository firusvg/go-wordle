package ansi

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nsf/termbox-go"

	tb "go-wordle/tbutility"
	"go-wordle/utility"
	"go-wordle/wordlist"
)

const (
	unusedLetter int = 0
	blackLetter  int = 1
	yellowLetter int = 2
	greenLetter  int = 3
)

type _letterStatus struct {
	letter rune
	x, y   int
	fg, bg termbox.Attribute
	status int
}

func createLetterStatusGrid(letters []rune, charsPerRow int, _x, _y int) []_letterStatus {
	var curs []rune
	result := make([]_letterStatus, len(letters))
	xPos := _x
	yPos := _y

	s, n := wordlist.SplitLetters(letters, charsPerRow)
	for i := 0; i < n; i++ {
		curs = []rune(s[i])
		if i == (n - 1) {
			xPos += (charsPerRow - len(curs)) / 2
		}
		for j := 0; j < len(curs); j++ {
			k := i*charsPerRow + j
			result[k].letter = rune(curs[j])
			result[k].x = xPos + j
			result[k].y = yPos + i
			result[k].fg = termbox.ColorBlack
			result[k].bg = termbox.ColorWhite
			result[k].status = unusedLetter
			termbox.SetCell(result[k].x, result[k].y, result[k].letter, result[k].fg, result[k].bg)
		}
	}
	return result
}

func indexLetter(slice *[]_letterStatus, c rune) int {
	for i := 0; i < len(*slice); i++ {
		if (*slice)[i].letter == c {
			return i
		}
	}
	return -1
}

func updateLetterStatusGrid(letterGrid *[]_letterStatus, _word, _wordStatus string) {
	word := []rune(_word)
	wordStatus := []rune(_wordStatus)

	for i := 0; i < len(wordStatus); i++ {
		j := indexLetter(letterGrid, word[i])
		letterStatus := (*letterGrid)[j]
		switch wordStatus[i] {
		case '🟨':
			letterStatus.fg = termbox.ColorBlack
			letterStatus.bg = termbox.ColorYellow
			letterStatus.status = yellowLetter
		case '🟩':
			letterStatus.fg = termbox.ColorBlack
			letterStatus.bg = termbox.ColorGreen
			letterStatus.status = greenLetter
		case '⬛':
			letterStatus.fg = termbox.ColorWhite
			letterStatus.bg = termbox.ColorBlack
			letterStatus.status = blackLetter
		}
		if letterStatus.status > (*letterGrid)[j].status {
			(*letterGrid)[j] = letterStatus
			termbox.SetCell((*letterGrid)[j].x, (*letterGrid)[j].y, (*letterGrid)[j].letter, (*letterGrid)[j].fg, (*letterGrid)[j].bg)
		}
	}
}

func Main() int {
	var finished bool
	var results [6]string
	var input string
	var escPressed bool
	var guessStatus int
	var resultString string
	var resultEmoji string
	var tries string

	wordleNo, wordleWord := wordlist.TodaysWordle()
	header := "Wordle " + strconv.Itoa(wordleNo)
	allowedRunes := wordlist.Letters
	tryNumber := 0
	curY := 3

	err := termbox.Init()
	if err != nil {
		return 1
	}

	letterStatus := createLetterStatusGrid(allowedRunes, 7, 18, 2)
	_, _ = tb.Print(1, 1, header, false)
	_, sizeY := termbox.Size()
	_, _ = tb.Print(1, sizeY-2, "Press ESC to quit.", false)
	termbox.SetCursor(1, curY)

	for !(finished) {
		input, escPressed = tb.Input(1, curY, 5, allowedRunes, true, ">")
		termbox.Flush()
		if escPressed {
			tb.ParkCursor(curY + 1)
			return 0
		} else {
			guessStatus, resultString, resultEmoji = utility.CheckWord(strings.ToUpper(input), wordleWord)
			_, _ = tb.Print(2, curY, resultString, true)
			if guessStatus < 0 {
				continue
			}
			updateLetterStatusGrid(&letterStatus, strings.ToUpper(input), resultEmoji)
			tryNumber++
			results[tryNumber-1] = resultEmoji
			switch guessStatus {
			case 0:
				finished = tryNumber == 6
			case 1:
				finished = true
			}
			curY++
		}
	}

	curY++
	if guessStatus == 0 {
		_, curY = tb.Print(1, curY, fmt.Sprintf("<r0>FAIL - %s<w0>", strings.ToUpper(wordleWord)), true)
		tries = "X/6"
	} else {
		_, curY = tb.Print(1, curY, fmt.Sprintf("<g0>WIN - %s<w0>", utility.WinStatuses[tryNumber-1]), true)
		tries = fmt.Sprintf("%d/6", tryNumber)
	}
	_, curY = tb.Print(0, curY+1, fmt.Sprintf("%s %s\n", header, tries), false)
	curY++
	for i := 0; i < tryNumber; i++ {
		_, curY = tb.Print(0, curY, results[i], false)
	}
	tb.ParkCursor(curY + 1)

	for {
		ev := termbox.PollEvent()
		switch ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				return 0
			}
		}
	}
}
