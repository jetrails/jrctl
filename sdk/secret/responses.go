package secret

type ErrorResponse struct {
	Name    string      `json:"name"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Type    string      `json:"type"`
	Data    interface{} `json:"data"`
}

type SecretCreateResponse struct {
	Identifier string `json:"id"`
	Password   string `json:"password"`
	TTL        int    `json:"ttl"`
}

type SecretReadResponse struct {
	Data string `json:"data"`
}

type SecretDeleteResponse struct {
	Identifier string `json:"id"`
}
