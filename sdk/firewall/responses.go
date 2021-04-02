package firewall

type ListResponse struct {
	Status string       `json:"status"`
	Code int            `json:"code"`
	Messages [] string  `json:"messages"`
	Payload [] Entry    `json:"payload"`
}

type AllowResponse struct {
	Status string       `json:"status"`
	Code int            `json:"code"`
	Messages [] string  `json:"messages"`
	Payload AllowRequest  `json:"payload"`
}

type CloudflareEntry struct {
	Address string            `json:"address"`
	Message string            `json:"message"`
}

type CloudflareEntries struct {
	Skipped [] string         `json:"skipped"`
	Allowed [] string         `json:"allowed"`
	Failed [] CloudflareEntry `json:"failed"`
}

type CloudflareResponse struct {
	Status string             `json:"status"`
	Code int                  `json:"code"`
	Messages [] string        `json:"messages"`
	Payload CloudflareEntries `json:"payload"`
}
