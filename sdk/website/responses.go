package website

import (
	"github.com/jetrails/jrctl/sdk/api"
)

type ListResponse struct {
	api.GenericResponse
	Payload []Properties `json:"payload"`
}

type PhpSwitchResponse struct {
	api.GenericResponse
	Payload interface{} `json:"payload"`
}

type PhpAvailableResponse struct {
	api.GenericResponse
	Payload []Availablity `json:"payload"`
}
