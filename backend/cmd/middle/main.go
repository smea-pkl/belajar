package main

import (
	"log/slog"
	"strings"
)

type ContentFunction func(content string)

func main() {
	content := "asdasdasdasda sbdasda saya sdasdas dasd aku"

	processContent(content)

}

func processContent(content string) {

	banyakFungsi := []ContentFunction{
		hitungA,
		hitungBadword,
		hitungKata,
	}

	for _, handler := range banyakFungsi {
		handler(content)
	}

}

func hitungA(content string) {
	count := 0

	for _, c := range content {
		switch c {
		case 'a', 'A':
			count++
		}
	}

	slog.Info("hitungA", "a", count)
}

func hitungKata(content string) {
	katas := strings.Split(content, "")
	jumlah := len(katas)

	slog.Info("hitungKata", "jumlah kata", jumlah)
}

func hitungBadword(content string) {
	lower := strings.ToLower(content)
	lowers := strings.Split(lower, " ")

	count := 0
	for _, k := range lowers {
		switch k {
		case "saya", "aku":
			count++
		}
	}

	slog.Info("hitungBadword", "jumlah badword", count)
}

func uniqueCharacter(content string) {
	unique := map[rune]bool{}
	for _, r := range content {
		unique[r] = true
	}

	jumlah := len(unique)
	slog.Info("uniqueCharacter", "jumlah unik", jumlah)
}
