package formatting

import (
	"fmt"
	"strings"
)

func GetColorSet() string {
	widestLen := getWidestLen()
	s := make([]string, 0, len(shortByLongColorName)+1)

	gap := strings.Repeat(" ", widestLen-4+2)
	s = append(s, fmt.Sprintf("wide%sshort", gap))

	for longName := range shortByLongColorName {
		gap = strings.Repeat(" ", widestLen-len(longName)+2)
		s = append(s, fmt.Sprintf(
			"%s%s%s%s%s",
			getSeq(longName),
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
	for longName := range shortByLongColorName {
		if len(longName) > l {
			l = len(longName)
		}
	}
	return l
}
