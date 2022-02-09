package utils

import "strings"

//CutString ...
func CutString(kind int, s string) (result string) {

	switch kind {
	case 1:
		result = s[76 : len(s)-4]
	case 2:
		result = s[75 : len(s)-4]
	case 3:
		result = s[89 : len(s)-4]
	case 4:
		if strings.Contains(s, "thumbnails") {
			return ""
		}
		result = s[79 : len(s)-4]
	}

	return
}
