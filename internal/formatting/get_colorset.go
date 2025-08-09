package formatting

import (
	"fmt"
	"strings"
)

func GetColorSet() string {
	widestLen := getWidestLen()
	s := make([]string, 0, len(colorSetWide)+1)

	gap := strings.Repeat(" ", widestLen-4+2)
	s = append(s, fmt.Sprintf("wide%sshort", gap))

	for longName, seq := range colorSetWide {
		gap = strings.Repeat(" ", widestLen-len(longName)+2)
		s = append(s, fmt.Sprintf(
			"%s%s%s%s%s",
			seq,
			longName,
			gap,
			shortByLongColorName[longName],
			clear,
		))
	}
	return strings.Join(s, "\n")
}

func getWidestLen() int {
	l := 0
	for longName, _ := range colorSetWide {
		if len(longName) > l {
			l = len(longName)
		}
	}
	return l
}
