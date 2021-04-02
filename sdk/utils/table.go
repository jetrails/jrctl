package utils

import (
	"fmt"
	"strings"
)

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
