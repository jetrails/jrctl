package service

type DaemonContext struct {
	Endpoint string  `json:"endpoint"`
	Debug bool       `json:"debug"`
	Auth string      `json:"auth"`
}
