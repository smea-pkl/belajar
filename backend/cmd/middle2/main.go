package main

import (
	"fmt"
	"log/slog"
	"slices"
	"strings"
	"unicode"
)

type ContentFunc func(content string) (string, error)
type ContentMiddleware func(next ContentFunc) ContentFunc

func main() {
	content := "saya"

	handler := createMiddleware([]ContentMiddleware{
		aDuplicate,
		func(next ContentFunc) ContentFunc {
			return func(content string) (string, error) {
				slog.Info("cegat", "content", content)

				if strings.Contains(content, "saya") {

					return "tidak diteruskan", nil
				}

				return next(content)
			}
		},
		countWord,
		encodeWord,
	})

	res, _ := handler(content)

	slog.Info("hasil", "hasil", res)

}

func createMiddleware(middlewares []ContentMiddleware) ContentFunc {

	var handler ContentFunc = func(content string) (string, error) {
		return content, nil
	}

	slices.Reverse(middlewares)
	for _, mid := range middlewares {
		handler = mid(handler)
	}

	return handler

}

func aDuplicate(next ContentFunc) ContentFunc {

	return func(content string) (string, error) {

		result := []string{}

		for _, r := range content {
			switch r {
			case 'a', 'A':
				result = append(result, content)
			}

		}

		return next(strings.Join(result, " "))
	}

}

func countWord(next ContentFunc) ContentFunc {

	return func(content string) (string, error) {
		slog.Info("running countWord")

		count := len(strings.Split(content, " "))
		return next(fmt.Sprintf("%d %s", count, content))
	}

}

func encodeWord(next ContentFunc) ContentFunc {

	return func(content string) (string, error) {

		var res string
		for _, c := range content {
			res += string(encodeRune(c))
		}

		return next(res)
	}

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
