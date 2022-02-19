package tbutility

import (
	"unicode"

	"go-wordle/utility"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

func Print(x, y int, msg string, parse bool) (int, int) {
	var fg, bg termbox.Attribute
	slice := []rune(msg)
	fg = termbox.ColorWhite
	bg = termbox.ColorDefault
	i := 0
	Colors := map[rune]termbox.Attribute{
		'Y': termbox.ColorYellow,
		'G': termbox.ColorGreen,
		'R': termbox.ColorRed,
		'W': termbox.ColorWhite,
		'B': termbox.ColorBlack,
		'0': termbox.ColorDefault,
	}

	for i < len(slice) {
		if parse && (slice[i] == '<') {
			fg = Colors[unicode.ToUpper(slice[i+1])]
			bg = Colors[unicode.ToUpper(slice[i+2])]
			i += 3
		} else {
			termbox.SetCell(x, y, slice[i], fg, bg)
			x += runewidth.RuneWidth(slice[i])
		}
		i++
	}

	return x, y + 1
}

func Input(x, y, maxlen int, allowedRunes []rune, ignoreCase bool, prompt string) (string, bool) {
	buffer := make([]rune, maxlen)
	p := 0

	_, _ = Print(x, y, prompt, false)
	x += len([]rune(prompt))
	for {
		termbox.SetCursor(x+p, y)
		termbox.Flush()
		ev := termbox.PollEvent()
		switch ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				return "", true
			} else if ev.Key == termbox.KeyEnter {
				if p == maxlen {
					return string(buffer), false
				}
			} else if (ev.Key == termbox.KeyBackspace) || (ev.Key == termbox.KeyBackspace2) {
				if p > 0 {
					p--
					buffer[p] = ' '
					termbox.SetCell(x+p, y, buffer[p], termbox.ColorWhite, termbox.ColorDefault)
				}
			} else {
				if p < maxlen {
					if utility.IndexRune(allowedRunes, unicode.ToUpper(ev.Ch)) >= 0 {
						buffer[p] = unicode.ToUpper(ev.Ch)
						termbox.SetCell(x+p, y, buffer[p], termbox.ColorWhite, termbox.ColorDefault)
						p++
					}
				}
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}

func ParkCursor(y int) {
	termbox.SetCursor(0, y)
	termbox.Flush()
}
