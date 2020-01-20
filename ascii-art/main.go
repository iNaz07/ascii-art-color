package main

import (
	"fmt"
	"os"

	u "./utils"
)

func main() {
	params := u.ReadParams(os.Args[1:])
	if params.Error != nil {
		fmt.Println(params.Error)
	} else {
		u.Printer(params)
	}
}
