package utils

import (
	"afr/hidisp"
	"strings"
)

func GetLine(text, str string) string {
	lines := strings.Split(text, "\n")
	for _, l := range lines {
		if strings.Contains(l, str) {
			return l
		}
	}
	return ""
}

func Is4K() bool {
	buf := hidisp.Hidisp("gethdmicap")
	line := GetLine(string(buf), "3840X2160_60")
	is4k := strings.Contains(line, "=1")
	return is4k
}
