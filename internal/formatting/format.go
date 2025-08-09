package formatting

import "fmt"

func Colorize(s, color string) string {
	if seq, ok := colorSetShort[color]; ok {
		return fmt.Sprintf("%s%s%s", seq, s, clear)
	}
	if seq, ok := colorSetWide[color]; ok {
		return fmt.Sprintf("%s%s%s", seq, s, clear)
	}
	return s
}
