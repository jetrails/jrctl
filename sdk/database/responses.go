package database

import (
	"github.com/jetrails/jrctl/sdk/api"
)

type ListDatabasesResponse struct {
	api.GenericResponse
	Payload []DatabaseWithUsers `json:"payload"`
}

type ListUsersResponse struct {
	api.GenericResponse
	Payload []UserWithDatabases `json:"payload"`
}

type CreateResponse struct {
	api.GenericResponse
	Payload interface{} `json:"payload"`
}

type DeleteResponse struct {
	api.GenericResponse
	Payload ConfirmPayload `json:"payload"`
}

type UserCreateResponse struct {
	api.GenericResponse
	Payload string `json:"payload"`
}

type UserDeleteResponse struct {
	api.GenericResponse
	Payload ConfirmPayload `json:"payload"`
}

type UserPasswordResponse struct {
	api.GenericResponse
	Payload ConfirmPayload `json:"payload"`
}

type LinkResponse struct {
	api.GenericResponse
	Payload interface{} `json:"payload"`
}

type UnlinkResponse struct {
	api.GenericResponse
	Payload ConfirmPayload `json:"payload"`
}
