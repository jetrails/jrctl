package confirm

import (
	"github.com/jetrails/jrctl/sdk/api"
)

type ConfirmResponse struct {
	api.GenericResponse
	Payload interface{} `json:"payload"`
}
