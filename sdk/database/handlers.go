package database

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jetrails/jrctl/sdk/api"
	"github.com/jetrails/jrctl/sdk/server"
	"github.com/jetrails/jrctl/sdk/version"
	"github.com/parnurzeal/gorequest"
)

func ListDatabases(context server.Context) ListDatabasesResponse {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	_, body, errs := request.
		Timeout(10*time.Second).
		Get(fmt.Sprintf("https://%s/database", context.Endpoint)).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Set("Authorization", "Bearer "+context.Token).
		Type("text").
		Send("{}").
		End()
	if len(errs) > 0 {
		return ListDatabasesResponse{
			GenericResponse: api.NewClientError(),
			Payload:         []DatabaseWithUsers{},
		}
	}
	var response ListDatabasesResponse
	json.Unmarshal([]byte(body), &response)
	return response
}

func ListUsers(context server.Context) ListUsersResponse {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	_, body, errs := request.
		Timeout(10*time.Second).
		Get(fmt.Sprintf("https://%s/database/user", context.Endpoint)).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Set("Authorization", "Bearer "+context.Token).
		Type("text").
		Send("{}").
		End()
	if len(errs) > 0 {
		return ListUsersResponse{
			GenericResponse: api.NewClientError(),
			Payload:         []UserWithDatabases{},
		}
	}
	var response ListUsersResponse
	json.Unmarshal([]byte(body), &response)
	return response
}

func Create(context server.Context, data CreateRequest) CreateResponse {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	_, body, errs := request.
		Timeout(10*time.Second).
		Put(fmt.Sprintf("https://%s/database", context.Endpoint)).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Set("Authorization", "Bearer "+context.Token).
		Send(data).
		End()
	if len(errs) > 0 {
		return CreateResponse{
			GenericResponse: api.NewClientError(),
			Payload:         "",
		}
	}
	var response CreateResponse
	json.Unmarshal([]byte(body), &response)
	return response
}

func Delete(context server.Context, data DeleteRequest) DeleteResponse {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	_, body, errs := request.
		Timeout(10*time.Second).
		Delete(fmt.Sprintf("https://%s/database", context.Endpoint)).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Set("Authorization", "Bearer "+context.Token).
		Send(data).
		End()
	if len(errs) > 0 {
		return DeleteResponse{
			GenericResponse: api.NewClientError(),
			Payload:         ConfirmPayload{},
		}
	}
	var response DeleteResponse
	json.Unmarshal([]byte(body), &response)
	return response
}

func UserCreate(context server.Context, data UserCreateRequest) UserCreateResponse {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	_, body, errs := request.
		Timeout(10*time.Second).
		Put(fmt.Sprintf("https://%s/database/user", context.Endpoint)).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Set("Authorization", "Bearer "+context.Token).
		Send(data).
		End()
	if len(errs) > 0 {
		return UserCreateResponse{
			GenericResponse: api.NewClientError(),
			Payload:         "",
		}
	}
	var response UserCreateResponse
	json.Unmarshal([]byte(body), &response)
	return response
}

func UserDelete(context server.Context, data UserDeleteRequest) UserDeleteResponse {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	_, body, errs := request.
		Timeout(10*time.Second).
		Delete(fmt.Sprintf("https://%s/database/user", context.Endpoint)).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Set("Authorization", "Bearer "+context.Token).
		Send(data).
		End()
	if len(errs) > 0 {
		return UserDeleteResponse{
			GenericResponse: api.NewClientError(),
			Payload:         ConfirmPayload{},
		}
	}
	var response UserDeleteResponse
	json.Unmarshal([]byte(body), &response)
	return response
}

func UserPassword(context server.Context, data UserPasswordRequest) UserPasswordResponse {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	_, body, errs := request.
		Timeout(10*time.Second).
		Put(fmt.Sprintf("https://%s/database/user/password", context.Endpoint)).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Set("Authorization", "Bearer "+context.Token).
		Send(data).
		End()
	if len(errs) > 0 {
		return UserPasswordResponse{
			GenericResponse: api.NewClientError(),
			Payload:         ConfirmPayload{},
		}
	}
	var response UserPasswordResponse
	json.Unmarshal([]byte(body), &response)
	return response
}

func Link(context server.Context, data LinkRequest) LinkResponse {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	_, body, errs := request.
		Timeout(10*time.Second).
		Put(fmt.Sprintf("https://%s/database/link", context.Endpoint)).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Set("Authorization", "Bearer "+context.Token).
		Send(data).
		End()
	if len(errs) > 0 {
		return LinkResponse{
			GenericResponse: api.NewClientError(),
			Payload:         ConfirmPayload{},
		}
	}
	var response LinkResponse
	json.Unmarshal([]byte(body), &response)
	return response
}

func Unlink(context server.Context, data UnlinkRequest) UnlinkResponse {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	_, body, errs := request.
		Timeout(10*time.Second).
		Delete(fmt.Sprintf("https://%s/database/link", context.Endpoint)).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Set("Authorization", "Bearer "+context.Token).
		Send(data).
		End()
	if len(errs) > 0 {
		return UnlinkResponse{
			GenericResponse: api.NewClientError(),
			Payload:         ConfirmPayload{},
		}
	}
	var response UnlinkResponse
	json.Unmarshal([]byte(body), &response)
	return response
}
