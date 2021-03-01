package service

type RestartResponse struct {
	Status string         `json:"status"`
	Code int              `json:"code"`
	Messages [] string    `json:"messages"`
	Payload interface {}  `json:"payload"`
}
