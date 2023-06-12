package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	separator := regexp.MustCompile(`<[~*=-]*>`)
	return separator.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	quotedPassword := regexp.MustCompile(`".*password.*"`)
	var count int
	for _, line := range lines {
		if quotedPassword.Match([]byte(strings.ToLower(line))) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	eol := regexp.MustCompile(`end-of-line\d+`)
	return eol.ReplaceAllString(text,"")
}

func TagWithUserName(lines []string) []string {
	user := regexp.MustCompile(`User\s+(\w+)`)
	outLines := []string(lines)
	for index, line := range outLines {
		m := user.FindStringSubmatch(line)
		if m != nil {
			outLines[index] = fmt.Sprintf("[USR] %s %s", m[1], line)
		}
	}
	return outLines
}
