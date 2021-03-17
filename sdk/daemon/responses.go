package daemon

type VersionResponse struct {
	Status string      `json:"status"`
	Code int           `json:"code"`
	Messages [] string `json:"messages"`
	Payload string     `json:"payload"`
}

type ListServicesResponse struct {
	Status string      `json:"status"`
	Code int           `json:"code"`
	Messages [] string `json:"messages"`
	Payload [] string  `json:"payload"`
}
