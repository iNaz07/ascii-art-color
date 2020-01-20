package main

import (
	"fmt"

	u "../utils"
)

func main() {

	var testing [][]string

	testing = append(testing, []string{"hello", "standard"})
	testing = append(testing, []string{"hello world", "shadow"})
	testing = append(testing, []string{"nice 2 meet you", "thinkertoy"})
	testing = append(testing, []string{"you & me", "standard"})
	testing = append(testing, []string{"123", "shadow"})
	testing = append(testing, []string{"/(\")", "thinkertoy"})
	testing = append(testing, []string{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", "shadow"})
	testing = append(testing, []string{"\"#$%&/()*+,-./", "thinkertoy"})
	testing = append(testing, []string{"It's Working", "thinkertoy"})

	for _, v := range testing {
		test(v)
	}
}

func test(p []string) {
	params := u.ReadParams(p)
	if params.Error != nil {
		fmt.Println(params.Error)
	} else {
		u.Printer(params)
	}
}
