package firewall

type Entry struct {
	Address string   `json:"address"`
	Port [] int      `json:"port"`
}
