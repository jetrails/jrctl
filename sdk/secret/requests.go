package secret

type SecretCreateRequest struct {
	Data string        `json:"data"`
	Password string    `json:"password"`
	TTL int            `json:"ttl,omitempty"`
	AutoGenerate bool  `json:"auto_generate"`
}

type SecretReadRequest struct {
	Identifier string  `json:"id"`
	Password string    `json:"password"`
}

type SecretDeleteRequest struct {
	Identifier string  `json:"id"`
}
