package report

import (
	"github.com/jetrails/jrctl/sdk/database"
	"github.com/jetrails/jrctl/sdk/firewall"
)

type AuditData struct {
	Whitelisted   []firewall.Entry             `json:"whitelisted"`
	PassAccess    []string                     `json:"pass_access"`
	KeyAccess     []string                     `json:"key_access"`
	Activity      []AccessLogEntry             `json:"activity"`
	Databases     []database.DatabaseWithUsers `json:"dbs"`
	DatabaseUsers []database.UserWithDatabases `json:"db_users"`
}

type AccessLogEntry struct {
	Month  string `json:"month"`
	Day    string `json:"day"`
	Time   string `json:"time"`
	Method string `json:"method"`
	User   string `json:"user"`
	Source string `json:"source"`
}
