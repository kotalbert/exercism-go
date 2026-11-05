package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)].*`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`".*pass[wW]ord.*"`)
	var c int
	for _, l := range lines {
		if re.MatchString(l) {
			c++
		}
	}
	return c
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\S+)`)
	for i, l := range lines {
		if re.MatchString(l) {
			u := re.FindStringSubmatch(l)[1]
			lines[i] = fmt.Sprintf("[USR] %s %s", u, l)
		}
	}
	return lines
}
