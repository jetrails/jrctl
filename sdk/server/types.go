package server

type Context struct {
	Endpoint string   `json:"endpoint"`
	Token    string   `json:"token"`
	Types    []string `json:"types"`
	Debug    bool     `json:"debug"`
	Insecure bool     `json:"insecure"`
}

type Entry struct {
	Endpoint string   `json:"endpoint"`
	Token    string   `json:"token"`
	Types    []string `json:"types"`
}

type TokenData struct {
	Identity         string   `json:"identity"`
	TokenID          string   `json:"token_id"`
	AllowedClientIPs []string `json:"allowed_client_ips"`
}

type ServiceProperties struct {
	Status  string `json:"status"`
	Restart bool   `json:"restart"`
	Reload  bool   `json:"reload"`
	Enable  bool   `json:"enable"`
	Disable bool   `json:"disable"`
}
