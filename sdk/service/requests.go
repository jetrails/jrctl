package service

type RestartRequest struct {
	Service string `json:"service"`
}

type ReloadRequest struct {
	Service string `json:"service"`
}

type EnableRequest struct {
	Service string `json:"service"`
}

type DisableRequest struct {
	Service string `json:"service"`
}
