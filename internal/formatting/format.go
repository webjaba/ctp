package formatting

import "fmt"

func Colorize(s, color string) string {
	return fmt.Sprintf("%s%s%s", getSeq(color), s, clear)
}

func getSeq(color string) string {
	if short, ok := shortByLongColorName[color]; ok {
		return colorSet[short]
	}
	return colorSet[color]
}
