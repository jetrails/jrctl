package firewall

type Entry struct {
	Action    string   `json:"action"`
	Source    string   `json:"source"`
	Ports     []int    `json:"ports"`
	Protocols []string `json:"protocols"`
	Comment   string   `json:"comment"`
}
