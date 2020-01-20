package utils

import "fmt"

func Printer(params Params) {
	fonts := GetFont(params.Font)

	for i := 1; i <= 8; i++ {
		if params.Range.Start == 0 && params.Range.End == 0 {
			for _, v := range params.Text {
				PrintColor(params.Color, fonts[int(v)-32][i])
			}
		} else {
			for j, v := range params.Text {
				if params.Range.Start <= j && params.Range.End > j {
					PrintColor(params.Color, fonts[int(v)-32][i])
				} else {
					PrintColor("white", fonts[int(v)-32][i])
				}
			}
		}
		fmt.Println()
	}

}
