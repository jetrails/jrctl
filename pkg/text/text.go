package text

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/jetrails/jrctl/pkg/env"
)

func Examples(examples []string) string {
	space := "  "
	if env.GetBool("docs", false) {
		space = ""
	}
	for i, e := range examples {
		examples[i] = space + e
	}
	return strings.Join(examples, "\n")
}

func Combine(lines []string) string {
	return strings.Join(lines, "\n\n")
}

func Paragraph(lines []string) string {
	width := 80
	result := ""
	line := ""
	delim := ""
	lineDelim := ""
	combined := strings.Join(lines, " ")
	for _, word := range strings.Split(combined, " ") {
		temp := line + delim + word
		if len(temp) <= width {
			line = temp
		} else {
			result = result + lineDelim + line
			line = word
			lineDelim = "\n"
		}
		delim = " "
	}
	return result + lineDelim + line
}

func TableString(message string, rows [][]string, padding int) string {
	if len(rows) < 2 {
		return strings.Repeat("\n", padding) + message + strings.Repeat("\n", padding+1)
	}
	output := strings.Repeat("\n", padding)
	max := []int{}
	for _, row := range rows {
		for c, col := range row {
			if len(max) <= c {
				max = append(max, len(col))
			}
			size := len(col)
			if max[c] < size {
				max[c] = size
			}
		}
	}
	for r, row := range rows {
		for c, col := range row {
			if r == 0 {
				col = strings.ToUpper(col)
			}
			if r > 0 && col == "" {
				col = "-"
			}
			format := fmt.Sprintf("%%-%ds  ", max[c])
			output += fmt.Sprintf(format, col)
		}
		output += "\n"
	}
	return output + strings.Repeat("\n", padding)
}

func TablePrint(message string, rows [][]string, padding int) {
	fmt.Print(TableString(message, rows, padding))
}

func SanitizeString(allowed, input string) string {
	return regexp.MustCompile(allowed).ReplaceAllString(input, "_")
}
