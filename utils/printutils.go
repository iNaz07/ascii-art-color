package utils

import (
	"fmt"
	"strings"
)

// Printer function
func Printer(fonts [][]string, params Params) {
	splittedText := strings.Split(params.Text, "\\n")
	for _, text := range splittedText {
		print(fonts, params.Color, text)
	}
}

func print(fonts [][]string, color []ColoredText, text string) {
	for i := 1; i <= 8; i++ {
		for _, v := range text {
			PrintColor(findColor(string(v), color).Color, fonts[int(v)-32][i])
		}
		fmt.Println()
	}
}

func findColor(s string, color []ColoredText) ColoredText {
	for _, c := range color {
		if c.Text == s {
			return c
		}
	}
	return ColoredText{
		Text:  "",
		Color: "white",
	}
}
