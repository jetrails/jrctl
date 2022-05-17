package server

type VersionResponse struct {
	Status   string   `json:"status"`
	Code     int      `json:"code"`
	Messages []string `json:"messages"`
	Payload  string   `json:"payload"`
}

type ListServicesResponse struct {
	Status   string                       `json:"status"`
	Code     int                          `json:"code"`
	Messages []string                     `json:"messages"`
	Metadata map[string]string            `json:"metadata"`
	Payload  map[string]ServiceProperties `json:"payload"`
}

type RestartResponse struct {
	Status   string      `json:"status"`
	Code     int         `json:"code"`
	Messages []string    `json:"messages"`
	Payload  interface{} `json:"payload"`
}

type ReloadResponse struct {
	Status   string      `json:"status"`
	Code     int         `json:"code"`
	Messages []string    `json:"messages"`
	Payload  interface{} `json:"payload"`
}

type EnableResponse struct {
	Status   string      `json:"status"`
	Code     int         `json:"code"`
	Messages []string    `json:"messages"`
	Payload  interface{} `json:"payload"`
}

type DisableResponse struct {
	Status   string      `json:"status"`
	Code     int         `json:"code"`
	Messages []string    `json:"messages"`
	Payload  interface{} `json:"payload"`
}

type TokenResponse struct {
	Status   string     `json:"status"`
	Code     int        `json:"code"`
	Messages []string   `json:"messages"`
	Payload  *TokenData `json:"payload"`
}
