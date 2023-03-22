package output

import (
	"fmt"
	"strings"

	"github.com/jetrails/jrctl/sdk/api"
	"github.com/jetrails/jrctl/sdk/config"
)

var ErrMsgNoServers = "No matching servers"
var ErrMsgNoResults = "No entries to display"
var ErrMsgRequiresOneServer = "multiple servers match, must narrow down to one server"

type Output struct {
	Servers           *Table
	Responses         []*Table
	Tags              []string
	Quiet             bool
	DisplayTags       bool
	DisplayServers    bool
	FailOnNoServers   bool
	FailOnNoResults   bool
	ExitCodeNoServers int
	ExitCodeNoResults int
	ErrMsgNoServers   Lines
	ErrMsgNoResults   Lines
}

func NewOutput(quiet bool, tags []string) *Output {
	servers := NewTable(Columns{
		"Hostname",
		"Server",
		"Tag(s)",
		"Response",
	})
	servers.Quiet = quiet
	return &Output{
		Servers:           servers,
		Responses:         []*Table{},
		Tags:              tags,
		Quiet:             quiet,
		DisplayTags:       true,
		DisplayServers:    false,
		FailOnNoServers:   false,
		FailOnNoResults:   false,
		ExitCodeNoServers: 0,
		ExitCodeNoResults: 0,
		ErrMsgNoServers:   Lines{ErrMsgNoServers},
		ErrMsgNoResults:   Lines{ErrMsgNoResults},
	}
}

func (o *Output) AddServer(context config.Context, response *api.GenericResponse, message string) {
	o.Servers.AddRow(Columns{
		response.Metadata["hostname"],
		strings.TrimSuffix(context.Endpoint, ":27482"),
		strings.Join(context.Tags, ", "),
		message,
	})
}

func (o *Output) AddUniqueServer(context config.Context, response *api.GenericResponse, message string) {
	for _, server := range o.Servers.Rows {
		if server[1] == strings.TrimSuffix(context.Endpoint, ":27482") {
			return
		}
	}
	o.Servers.AddRow(Columns{
		response.Metadata["hostname"],
		strings.TrimSuffix(context.Endpoint, ":27482"),
		strings.Join(context.Tags, ", "),
		message,
	})
}

func (o *Output) AddTable(tbl *Table) {
	tbl.Quiet = o.Quiet
	tbl.EmptyMessage = Lines{}
	o.Responses = append(o.Responses, tbl)
}

func (o *Output) CreateTable(cols Columns) *Table {
	tbl := NewTable(cols)
	o.AddTable(tbl)
	return tbl
}

func (o *Output) IsResultsEmpty() bool {
	for _, tbl := range o.Responses {
		if !tbl.IsEmpty() {
			return false
		}
	}
	return true
}

func (o *Output) GetString() string {
	result := ""
	result += o.GetDivider()
	if o.DisplayServers {
		result += o.Servers.GetString()
		result += o.GetDivider()
	}
	for _, tbl := range o.Responses {
		result += tbl.GetString()
		if !tbl.IsEmpty() {
			result += o.GetDivider()
		}
	}
	return result
}

func (o *Output) GetTagsString() string {
	result := ""
	or := ""
	if len(o.Tags) == 0 {
		return "ANY"
	}
	for _, tag := range o.Tags {
		parts := strings.Split(tag, ",")
		filtered := []string{}
		for _, part := range parts {
			trimmed := strings.TrimSpace(part)
			if trimmed != "" {
				filtered = append(filtered, trimmed)
			}
		}
		result += or
		or = " or "
		if len(filtered) == 0 {
			return "ANY"
		}
		if len(filtered) == 1 {
			result += filtered[0]
		} else {
			result += fmt.Sprintf("(%s)", strings.Join(filtered, " and "))
		}
	}
	return result
}

func (o *Output) PrintTags() {
	if !o.Quiet && o.DisplayTags {
		fmt.Print(o.GetDivider())
		fmt.Printf("Filtering servers that match the following type(s): ")
		fmt.Printf(o.GetTagsString())
		fmt.Printf(o.GetDivider())
	}
}

func (o *Output) Print() {
	divider := o.GetDivider()
	o.PrintTags()
	if o.FailOnNoServers && o.Servers.IsEmpty() {
		fmt.Print(divider)
		o.ExitWithMessage(o.ExitCodeNoServers, strings.Join(o.ErrMsgNoServers, divider)+divider)
	}
	fmt.Print(o.GetString())
	if o.FailOnNoResults && o.IsResultsEmpty() && len(o.ErrMsgNoResults) > 0 {
		if o.Servers.IsEmpty() && o.DisplayServers {
			fmt.Print(divider)
		}
		o.ExitWithMessage(o.ExitCodeNoResults, strings.Join(o.ErrMsgNoResults, divider)+divider)
	}
}

func (o *Output) GetDivider() string {
	return o.Servers.GetDivider()
}

func (o *Output) PrintDivider() {
	o.Servers.PrintDivider()
}

func (o *Output) ExitWithMessage(code int, format string, args ...interface{}) {
	o.Servers.ExitWithMessage(code, format, args...)
}

func (o *Output) PrintResponse(res *api.GenericResponse) {
	o.Servers.PrintResponse(res)
}

func (o *Output) ExitCodeFromResponse(res *api.GenericResponse) {
	o.Servers.ExitCodeFromResponse(res)
}
