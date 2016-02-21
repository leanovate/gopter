package gopter

import "strings"

func formatLines(str, lead, trail string, width int) string {
	result := ""
	for _, line := range strings.Split(str, "\n") {
		if result != "" {
			result += "\n"
		}
		result += breakLine(lead+line+trail, "  ", width)
	}
	return result
}

func breakLine(str, lead string, length int) string {
	result := ""
	for len(str) > length {
		result += lead + str[0:length] + "\n"
		str = str[length:]
	}
	result += str
	return result
}

func concatLines(strs ...string) string {
	result := ""
	for _, str := range strs {
		if str != "" {
			if result != "" {
				result += "\n"
			}
			result += str
		}
	}
	return result
}
