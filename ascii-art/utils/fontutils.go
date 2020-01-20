package utils

import (
	"bufio"
	"fmt"
	"os"
)

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
