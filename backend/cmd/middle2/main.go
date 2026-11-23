package main

import (
	"fmt"
	"log/slog"
	"strings"
	"unicode"
)

func main() {
	content := "sSAYA adalaH superhero"

	content = aDuplicate(content)
	content = countWord(content)
	content = encodeWord(content)

	slog.Info(content)

}

func aDuplicate(content string) string {
	result := []string{}

	for _, r := range content {
		switch r {
		case 'a', 'A':
			result = append(result, content)
		}

	}

	return strings.Join(result, " ")
}

func countWord(content string) string {
	count := len(strings.Split(content, " "))
	return fmt.Sprintf("%d %s", count, content)
}

func encodeWord(content string) string {
	var res string
	for _, c := range content {
		res += string(encodeRune(c))
	}

	return res
}

func encodeRune(c rune) rune {
	alfa := "01234567890abcdefghijklmnopqrstuvwxyza"

	var i int
	var ch rune
	var iscapital bool

	if c == ' ' {
		return c
	}

	for i, ch = range alfa {

		if c == ch {
			break
		}
		if unicode.ToLower(c) == ch {
			iscapital = true
			break
		}
	}

	chares := alfa[i+1]

	if iscapital {
		return unicode.ToUpper(rune(chares))
	}

	return rune(chares)

}
