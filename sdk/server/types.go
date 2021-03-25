package server

type Context struct {
	Endpoint string  `json:"endpoint"`
	Token string     `json:"token"`
	Types [] string  `json:"types"`
	Debug bool       `json:"debug"`
	Insecure bool    `json:"insecure"`
}

type Entry struct {
	Endpoint string  `json:"endpoint"`
	Token string     `json:"token"`
	Types [] string  `json:"types"`
}
