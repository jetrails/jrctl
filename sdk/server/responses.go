package server

type VersionResponse struct {
	Status   string   `json:"status"`
	Code     int      `json:"code"`
	Messages []string `json:"messages"`
	Payload  string   `json:"payload"`
}

type ListServicesResponse struct {
	Status   string   `json:"status"`
	Code     int      `json:"code"`
	Messages []string `json:"messages"`
	Payload  []string `json:"payload"`
}

type RestartResponse struct {
	Status   string      `json:"status"`
	Code     int         `json:"code"`
	Messages []string    `json:"messages"`
	Payload  interface{} `json:"payload"`
}
