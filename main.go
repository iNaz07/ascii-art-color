package main

import (
	"fmt"
	"os"

	u "./utils"
)

func main() {
	params := u.ReadParams(os.Args[1:])
	if params.Error != "" {
		fmt.Println(params.Error)
	} else {
		fonts := u.GetFont(params.Font)

		if params.ReversePath == "" {
			u.Printer(fonts, params)
		} else {
			arr := u.GetFile(params.ReversePath)
			reverse(fonts, arr)
		}
	}
}

func reverse(fonts [][]string, arr []string) {
	str := ""
	prev := 0

	for col := 0; col < len(arr[0]); col++ {
		found := true
		for row := 0; row < 8; row++ {
			if arr[row][col] != ' ' {
				found = false
			}
		}
		if found {
			str += u.GetLetter(fonts, u.Cut(arr, prev, col))
			prev = col + 1
		}
	}
	fmt.Println(str)
}
