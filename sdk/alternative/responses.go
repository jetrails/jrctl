package alternative

import (
	"github.com/jetrails/jrctl/sdk/api"
)

type ListResponse struct {
	api.GenericResponse
	Payload []Entry `json:"payload"`
}

type SwitchResponse struct {
	api.GenericResponse
	Payload interface{} `json:"payload"`
}
