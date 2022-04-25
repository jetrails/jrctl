package report

import (
	"github.com/jetrails/jrctl/sdk/firewall"
)

type AuditData struct {
	Whitelisted []firewall.Entry `json:"whitelisted"`
	PassAccess  []string         `json:"pass_access"`
	KeyAccess   []string         `json:"key_access"`
	Activity    []AccessLogEntry `json:"activity"`
}

type AccessLogEntry struct {
	Month  string `json:"month"`
	Day    string `json:"day"`
	Time   string `json:"time"`
	Method string `json:"method"`
	User   string `json:"user"`
	Source string `json:"source"`
}
