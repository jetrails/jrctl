package service

type ServiceProperties struct {
	Status  string `json:"status"`
	Restart bool   `json:"restart"`
	Reload  bool   `json:"reload"`
	Enable  bool   `json:"enable"`
	Disable bool   `json:"disable"`
}
