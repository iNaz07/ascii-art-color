package utils

import (
	"bufio"
	"fmt"
	"os"
)

// GetFont function
func GetFont(filename string) [][]string {
	fonts := [][]string{}

	file, err := os.Open("./fonts/" + filename + ".txt")

	if err != nil {
		fmt.Println("such a file does not exist")
	}
	defer file.Close()

	char := []string{}
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if i < 8 {
			char = append(char, scanner.Text())
			i++
		} else {
			char = append(char, scanner.Text())
			fonts = append(fonts, char)
			i = 0
			char = []string{}
		}
	}
	if err := scanner.Err(); err != nil {
	}
	return fonts
}

// GetLetter function
func GetLetter(fonts [][]string, letter []string) string {
	for i := 0; i < len(fonts); i++ {
		if compareFont(fonts[i], letter) {
			return string(i + 32)
		}
	}
	return ""
}

func compareFont(font, letter []string) bool {
	for col := 0; col < len(letter[1]); col++ {
		for row := 1; row < 8; row++ {
			if len(font[row]) >= len(letter[row]) {
				if font[row][col] != letter[row][col] {
					return false
				}
			} else {
				return false
			}
		}
	}
	return true
}

func Cut(arr []string, start, end int) []string {
	letter := []string{}
	letter = append(letter, "")
	for row := 0; row < 8; row++ {
		str := ""
		for col := start; col <= end; col++ {
			str += string(arr[row][col])
		}
		letter = append(letter, str)
	}
	return letter
}

func GetFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("such a file does not exist")
	}
	defer file.Close()

	arr := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}
	return arr
}
