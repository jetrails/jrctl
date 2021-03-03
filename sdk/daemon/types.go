package daemon

type Context struct {
	Debug bool       `json:"debug"`
	Endpoint string  `json:"endpoint"`
	Token string     `json:"token"`
	Tags [] string   `json:"tags"`
}

type DaemonConfig struct {
	Auth string  `yaml:"auth"`
}
