package firewall

import (
	"github.com/jetrails/jrctl/sdk/api"
)

type ListResponse struct {
	api.GenericResponse
	Payload []Entry `json:"payload"`
}

type AllowResponse struct {
	api.GenericResponse
	Payload AllowRequest `json:"payload"`
}

type UnAllowResponse struct {
	api.GenericResponse
	Payload UnAllowRequest `json:"payload"`
}

type DenyResponse struct {
	api.GenericResponse
	Payload DenyRequest `json:"payload"`
}

type UnDenyResponse struct {
	api.GenericResponse
	Payload UnDenyRequest `json:"payload"`
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
	api.GenericResponse
	Payload CloudflareEntries `json:"payload"`
}

type UnAllowCloudflareResponse struct {
	api.GenericResponse
	Payload CloudflareEntries `json:"payload"`
}
