package formatting

var (
	clear = "\033[0m"

	colorSet = map[string]string{
		"r":   "\033[01;38;05;196m",
		"g":   "\033[32m",
		"y":   "\033[01;38;05;226m",
		"b":   "\033[34m",
		"p":   "\033[35m",
		"c":   "\033[36m",
		"brp": "\033[01;38;05;201m",
		"lp":  "\033[01;38;05;217m",
	}

	shortByLongColorName = map[string]string{
		"green":       "g",
		"red":         "r",
		"bold-red":    "br",
		"yellow":      "y",
		"blue":        "b",
		"purple":      "p",
		"cyan":        "c",
		"bright-pink": "brp",
		"light-pink":  "lp",
	}
)
