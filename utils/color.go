package utils

import (
	"fmt"
	"regexp"
)

func PrintColor(color, text string) {
	colors := map[string]string{}
	colors["grey"] = "\033[1;30m%s\033[0m"
	colors["red"] = "\033[1;31m%s\033[0m"
	colors["green"] = "\033[1;32m%s\033[0m"
	colors["yellow"] = "\033[1;33m%s\033[0m"
	colors["blue"] = "\033[1;34m%s\033[0m"
	colors["magenta"] = "\033[1;35m%s\033[0m"
	colors["teal"] = "\033[1;36m%s\033[0m"
	colors["white"] = "\033[1;37m%s\033[0m"

	if color[:3] == "rgb" {
		r := regexp.MustCompile(`(?m)rgb\((\d{1,3}),(\d{1,3}),(\d{1,3})\)`)
		rgb := r.FindStringSubmatch(color)
		fmt.Printf("\x1b[38;2;%v;%v;%vm%s\x1b[0m", rgb[1], rgb[2], rgb[3], text)
	} else {
		fmt.Printf(colors[color], text)
	}
}
