package utils

type ColoredText struct {
	Text  string
	Color string
}

type Params struct {
	Text        string
	Font        string
	Color       []ColoredText
	ReversePath string
	Error       string
}

func ReadParams(args []string) Params {

	fontTypes := map[string]bool{}
	fontTypes["standard"] = true
	fontTypes["shadow"] = true
	fontTypes["thinkertoy"] = true

	params := Params{
		Font:        "standard",
		Text:        "",
		ReversePath: "",
	}

	fontFounded := false

	for _, v := range args {
		if fontTypes[v] && !fontFounded {
			params.Font = v
			fontFounded = true
		} else if len(v) > 8 && v[:8] == "--color=" {
			params.Color = append(params.Color, getColoredText(v[8:]))
		} else if len(v) > 10 && v[:10] == "--reverse=" {
			params.ReversePath = v[10:]
		} else {
			if params.Text != "" {
				params.Error = "Invalid font name"
			} else {
				params.Text = v
			}
		}
	}

	return params
}

func getColoredText(s string) ColoredText {
	coloredText := ColoredText{
		Text:  "",
		Color: "",
	}
	next := false
	for _, v := range s {
		if v == ':' {
			next = true
		} else if !next {
			coloredText.Color += string(v)
		} else {
			coloredText.Text += string(v)
		}
	}

	if coloredText.Color == "" {
		coloredText.Color = "white"
	}

	return coloredText
}
