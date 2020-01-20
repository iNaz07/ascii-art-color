package utils

type RangeStruct struct {
	Start int
	End   int
}

type Params struct {
	Text  string
	Font  string
	Color string
	Range RangeStruct
	Error *string
}

func ReadParams(args []string) Params {

	fontTypes := map[string]bool{}
	fontTypes["standard"] = true
	fontTypes["shadow"] = true
	fontTypes["thinkertoy"] = true

	params := Params{
		Font:  "standard",
		Text:  "",
		Color: "white",
	}

	for _, v := range args {
		if fontTypes[v] {
			params.Font = v
		} else if len(v) > 8 && v[:8] == "--color=" {
			params.Color = v[8:]
		} else if len(v) > 8 && v[:8] == "--range=" {
			params.Range = getRange(v[8:])
		} else {
			params.Text = v
		}
	}

	return params
}

func getRange(s string) RangeStruct {
	r := RangeStruct{
		Start: 0,
		End:   0,
	}

	isStart := true
	for _, v := range s {
		if v == ':' {
			isStart = false
		} else if isStart {
			if v >= 48 && v <= 57 {
				r.Start = r.Start*10 + int(v) - 48
			}
		} else {
			if v >= 48 && v <= 57 {
				r.End = r.End*10 + int(v) - 48
			}
		}
	}

	return r
}
