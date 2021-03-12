package service

type RestartRequest struct {
	Service string  `json:"service"`
	Version string  `json:"version"`
}
