package whitelist

type ListResponse struct {
	Status string       `json:"status"`
	Code int            `json:"code"`
	Messages [] string  `json:"messages"`
	Payload [] Entry    `json:"payload"`
}

type AddResponse struct {
	Status string       `json:"status"`
	Code int            `json:"code"`
	Messages [] string  `json:"messages"`
	Payload AddRequest  `json:"payload"`
}
