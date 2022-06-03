package website

import (
	"github.com/jetrails/jrctl/sdk/api"
)

type ListResponse struct {
	api.GenericResponse
	Payload []Properties `json:"payload"`
}

type SwitchPHPResponse struct {
	api.GenericResponse
	Payload interface{} `json:"payload"`
}
