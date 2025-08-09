package formatting

var (
	clear = "\033[0m"

	colorSetShort = map[string]string{
		"g":  "\033[32m",
		"r":  "\033[31m",
		"br": "\033[1;31m",
	}

	colorSetWide = map[string]string{
		"green":    "\033[32m",
		"red":      "\033[31m",
		"bold-red": "\033[1;31m",
	}

	shortByLongColorName = map[string]string{
		"green":    "g",
		"red":      "r",
		"bold-red": "br",
	}
)
