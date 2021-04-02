package secret

type PublicApiContext struct {
	Endpoint string `json:"endpoint"`
	Debug    bool   `json:"debug"`
	Insecure bool   `json:"insecure"`
}
