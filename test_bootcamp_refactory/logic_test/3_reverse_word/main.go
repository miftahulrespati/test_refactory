package main

import (
	"fmt"
	"strings"
	"unicode"
)

func rev_letters(s string) string {
	ltr := []rune(s)
	for i, j := 0, len(ltr)-1; i < j; i, j = i+1, j-1 {
		ltr[i], ltr[j] = ltr[j], ltr[i]
	}
	return string(ltr)
}

func rev_words(s string) string {
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func uppercase(str string, rts string) string {
	ltr := []rune(str)
	rtl := []rune(rts)
	for i := 0; i < len(str); i++ {
		rtl[i] = unicode.ToLower(rtl[i])
		if unicode.IsUpper(ltr[i]) {
			rtl[i] = unicode.ToUpper(rtl[i])
		}
	}
	return string(rtl)
}

func main() {
	str := "I am A Great human"
	strReverse := rev_letters(str)
	strReverse = rev_words(strReverse)
	strReverse = uppercase(str, strReverse)
	fmt.Println(strReverse)
}
