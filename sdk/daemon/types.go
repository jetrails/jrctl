package daemon

type Context struct {
	Debug bool         `json:"debug"`
	Nickname string    `json:"nickname"`
	Endpoint string    `json:"endpoint"`
	Token string       `json:"token"`
	Services [] string `json:"services"`
}
