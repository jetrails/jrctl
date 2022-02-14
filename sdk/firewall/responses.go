package firewall

type ListResponse struct {
	Status   string            `json:"status"`
	Code     int               `json:"code"`
	Messages []string          `json:"messages"`
	Metadata map[string]string `json:"metadata"`
	Payload  []Entry           `json:"payload"`
}

type AllowResponse struct {
	Status   string       `json:"status"`
	Code     int          `json:"code"`
	Messages []string     `json:"messages"`
	Payload  AllowRequest `json:"payload"`
}

type UnAllowResponse struct {
	Status   string         `json:"status"`
	Code     int            `json:"code"`
	Messages []string       `json:"messages"`
	Payload  UnAllowRequest `json:"payload"`
}

type DenyResponse struct {
	Status   string      `json:"status"`
	Code     int         `json:"code"`
	Messages []string    `json:"messages"`
	Payload  DenyRequest `json:"payload"`
}

type UnDenyResponse struct {
	Status   string        `json:"status"`
	Code     int           `json:"code"`
	Messages []string      `json:"messages"`
	Payload  UnDenyRequest `json:"payload"`
}

type CloudflareEntry struct {
	Address string `json:"address"`
	Message string `json:"message"`
}

type CloudflareEntries struct {
	Skipped   []string          `json:"skipped"`
	Succeeded []string          `json:"succeeded"`
	Failed    []CloudflareEntry `json:"failed"`
}

type AllowCloudflareResponse struct {
	Status   string            `json:"status"`
	Code     int               `json:"code"`
	Messages []string          `json:"messages"`
	Payload  CloudflareEntries `json:"payload"`
}

type UnAllowCloudflareResponse struct {
	Status   string            `json:"status"`
	Code     int               `json:"code"`
	Messages []string          `json:"messages"`
	Payload  CloudflareEntries `json:"payload"`
}
