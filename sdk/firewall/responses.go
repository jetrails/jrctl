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
