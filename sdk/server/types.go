package server

type Context struct {
	Debug bool       `json:"debug"`
	Endpoint string  `json:"endpoint"`
	Token string     `json:"token"`
	Types [] string  `json:"types"`
}

type ServerConfig struct {
	Auth string  `yaml:"auth"`
}
