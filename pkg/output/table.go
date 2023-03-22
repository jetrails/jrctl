package output

import (
	"crypto/sha256"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/jetrails/jrctl/pkg/array"
	"github.com/jetrails/jrctl/sdk/api"
)

type Lines []string

type Columns []string

type Table struct {
	Title           Lines
	Header          Columns
	Rows            []Columns
	EmptyMessage    Lines
	Quiet           bool
	QuietCollection []string
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func NewTable(header Columns) *Table {
	return &Table{
		Title:           Lines{},
		Header:          header,
		Rows:            []Columns{},
		EmptyMessage:    Lines{"No entries found"},
		Quiet:           false,
		QuietCollection: []string{},
	}
}

func (cols Columns) Clone() Columns {
	result := Columns{}
	for _, val := range cols {
		result = append(result, val)
	}
	return result
}

func (tbl *Table) AddRow(row Columns) {
	tbl.Rows = append(tbl.Rows, row)
}

func (tbl *Table) AddQuietEntry(entry string) {
	tbl.QuietCollection = append(tbl.QuietCollection, entry)
}

func (tbl *Table) AddUniqueQuietEntry(entry string) {
	if !array.ContainsString(tbl.QuietCollection, entry) {
		tbl.QuietCollection = append(tbl.QuietCollection, entry)
	}
}

func (tbl *Table) Sort(index int) {
	sort.SliceStable(tbl.Rows, func(i, j int) bool {
		return tbl.Rows[i][index] < tbl.Rows[j][index]
	})
}

func (tbl *Table) SortQuiet() {
	sort.Strings(tbl.QuietCollection)
}

func (tbl *Table) GetString() string {
	output := ""
	if tbl.Quiet {
		if len(tbl.QuietCollection) > 0 {
			tbl.SortQuiet()
			output += strings.Join(tbl.QuietCollection, "\n") + "\n"
		}
		return output
	}
	if len(tbl.Title) > 0 {
		output += "#\n"
		for _, line := range tbl.Title {
			output += "# " + line + "\n"
		}
		output += "#\n\n"
	}
	if len(tbl.Rows) < 1 {
		if len(tbl.EmptyMessage) > 0 {
			output += strings.Join(tbl.EmptyMessage, "\n") + "\n\n"
		}
		return output
	}
	max := []int{}
	maxCols := 0
	for _, row := range append(tbl.Rows, tbl.Header) {
		if len(row) > maxCols {
			maxCols = len(row)
		}
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
	for c := 0; c < maxCols; c++ {
		col := "-"
		if c < len(tbl.Header) && tbl.Header[c] != "" {
			col = tbl.Header[c]
			col = strings.ToUpper(col)
		}
		format := fmt.Sprintf("%%-%ds  ", max[c])
		output += fmt.Sprintf(format, col)
	}
	output += "\n"
	for _, row := range tbl.Rows {
		numCols := len(row)
		for c := 0; c < maxCols; c++ {
			col := "-"
			if c < numCols && row[c] != "" {
				col = row[c]
			}
			format := fmt.Sprintf("%%-%ds  ", max[c])
			output += fmt.Sprintf(format, col)
		}
		output += "\n"
	}
	return output
}

func (tbl *Table) PrintTable() {
	fmt.Print(tbl.GetString())
}

func (tbl *Table) IsEmpty() bool {
	return len(tbl.Rows) < 1
}

func (tbl *Table) SquashOnPivot(pivot int) {
	squashed := []Columns{}
	cache := map[string]Columns{}
	for _, row := range tbl.Rows {
		cloned := row.Clone()
		similar := append(cloned[0:pivot], cloned[pivot+1:]...)
		hash := sha256.New()
		hash.Write([]byte(fmt.Sprintf("%v", similar)))
		key := fmt.Sprintf("%x", hash.Sum(nil))
		if _, ok := cache[key]; ok {
			cache[key][pivot] += ", " + row[pivot]
		} else {
			cache[key] = row
		}
	}
	for _, row := range cache {
		squashed = append(squashed, row)
	}
	tbl.Rows = squashed
}

func (tbl *Table) GetDivider() string {
	if tbl.Quiet {
		return ""
	}
	return "\n"
}

func (tbl *Table) PrintDivider() {
	fmt.Print(tbl.GetDivider())
}

func (tbl *Table) ExitWithMessage(code int, format string, args ...interface{}) {
	if !tbl.Quiet {
		fmt.Printf(format+"\n", args...)
	}
	os.Exit(code)
}

func (tbl *Table) PrintResponse(res *api.GenericResponse) {
	if !tbl.Quiet {
		tbl := NewTable(Columns{
			"Status",
			"Message",
		})
		for _, message := range res.Messages {
			tbl.AddRow(Columns{
				fmt.Sprintf("%s (%d)", res.Status, res.Code),
				message,
			})
		}
		if len(res.Messages) == 0 {
			tbl.AddRow(Columns{
				fmt.Sprintf("%s (%d)", res.Status, res.Code),
			})
		}
		tbl.PrintTable()
	}
}

func (tbl *Table) ExitCodeFromResponse(res *api.GenericResponse) {
	if res.IsOkay() {
		os.Exit(0)
	}
	switch res.Code {
	// 1 is reserved for "No matching nodes"
	// 2 is reserved for "No results"
	case 3:
		os.Exit(3) // Reserved for "Client-Side Error"
	case 200:
		os.Exit(0) // http.StatusOK
	case 400:
		os.Exit(20) // http.StatusBadRequest
	case 401:
		os.Exit(21) // http.StatusUnauthorized
	case 403:
		os.Exit(22) // http.StatusForbidden
	case 404:
		os.Exit(23) // http.StatusNotFound
	case 405:
		os.Exit(24) // http.StatusMethodNotAllowed
	case 409:
		os.Exit(25) // http.StatusConflict
	case 422:
		os.Exit(26) // http.StatusUnprocessableEntity
	case 500:
		os.Exit(27) // http.StatusInternalServerError
	case 503:
		os.Exit(28) // http.StatusServiceUnavailable
	default:
		os.Exit(4) // Reserved for "Other Unkown Error"
	}
}
