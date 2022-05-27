package server

import (
	"github.com/jetrails/jrctl/sdk/api"
)

type VersionResponse struct {
	api.GenericResponse
	Payload string `json:"payload"`
}

type TokenResponse struct {
	api.GenericResponse
	Payload *TokenData `json:"payload"`
}
