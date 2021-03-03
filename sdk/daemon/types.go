package daemon

type Context struct {
	Debug bool         `json:"debug"`
	Endpoint string    `json:"endpoint"`
	Token string       `json:"token"`
	Services [] string `json:"services"`
}

type DaemonConfig struct {
	Auth string `yaml:"auth"`
}
