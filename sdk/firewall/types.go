package firewall

type Entry struct {
	Address string   `json:"address"`
	Port [] int      `json:"port"`
}

type DaemonContext struct {
	Endpoint string  `json:"endpoint"`
	Debug bool       `json:"debug"`
	Auth string      `json:"auth"`
}
