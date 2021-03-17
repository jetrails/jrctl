package daemon

type Context struct {
	Debug bool       `json:"debug"`
	Endpoint string  `json:"endpoint"`
	Token string     `json:"token"`
	Types [] string  `json:"types"`
}

type DaemonConfig struct {
	Auth string  `yaml:"auth"`
}
