package service

import (
	"github.com/jetrails/jrctl/sdk/api"
)

type ListServicesResponse struct {
	api.GenericResponse
	Payload map[string]ServiceProperties `json:"payload"`
}

type RestartResponse struct {
	api.GenericResponse
	Payload interface{} `json:"payload"`
}

type ReloadResponse struct {
	api.GenericResponse
	Payload interface{} `json:"payload"`
}

type EnableResponse struct {
	api.GenericResponse
	Payload interface{} `json:"payload"`
}

type DisableResponse struct {
	api.GenericResponse
	Payload interface{} `json:"payload"`
}
