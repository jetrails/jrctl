package database

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jetrails/jrctl/sdk/api"
	"github.com/jetrails/jrctl/sdk/config"
	"github.com/jetrails/jrctl/sdk/version"
	"github.com/parnurzeal/gorequest"
)

func ListDatabases(context config.Context) ListDatabasesResponse {
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
		if context.Debug {
			fmt.Println("Client Request Error:")
			fmt.Println(errs)
		}
		return ListDatabasesResponse{
			GenericResponse: api.NewClientError(),
			Payload:         []DatabaseWithUsers{},
		}
	}
	var response ListDatabasesResponse
	json.Unmarshal([]byte(body), &response)
	return response
}

func ListUsers(context config.Context) ListUsersResponse {
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
		if context.Debug {
			fmt.Println("Client Request Error:")
			fmt.Println(errs)
		}
		return ListUsersResponse{
			GenericResponse: api.NewClientError(),
			Payload:         []UserWithDatabases{},
		}
	}
	var response ListUsersResponse
	json.Unmarshal([]byte(body), &response)
	return response
}

func Create(context config.Context, data CreateRequest) CreateResponse {
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
		if context.Debug {
			fmt.Println("Client Request Error:")
			fmt.Println(errs)
		}
		return CreateResponse{
			GenericResponse: api.NewClientError(),
			Payload:         "",
		}
	}
	var response CreateResponse
	json.Unmarshal([]byte(body), &response)
	return response
}

func Delete(context config.Context, data DeleteRequest) DeleteResponse {
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
		if context.Debug {
			fmt.Println("Client Request Error:")
			fmt.Println(errs)
		}
		return DeleteResponse{
			GenericResponse: api.NewClientError(),
			Payload:         ConfirmPayload{},
		}
	}
	var response DeleteResponse
	json.Unmarshal([]byte(body), &response)
	return response
}

func UserCreate(context config.Context, data UserCreateRequest) UserCreateResponse {
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
		if context.Debug {
			fmt.Println("Client Request Error:")
			fmt.Println(errs)
		}
		return UserCreateResponse{
			GenericResponse: api.NewClientError(),
			Payload:         "",
		}
	}
	var response UserCreateResponse
	json.Unmarshal([]byte(body), &response)
	return response
}

func UserDelete(context config.Context, data UserDeleteRequest) UserDeleteResponse {
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
		if context.Debug {
			fmt.Println("Client Request Error:")
			fmt.Println(errs)
		}
		return UserDeleteResponse{
			GenericResponse: api.NewClientError(),
			Payload:         ConfirmPayload{},
		}
	}
	var response UserDeleteResponse
	json.Unmarshal([]byte(body), &response)
	return response
}

func UserPassword(context config.Context, data UserPasswordRequest) UserPasswordResponse {
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
		if context.Debug {
			fmt.Println("Client Request Error:")
			fmt.Println(errs)
		}
		return UserPasswordResponse{
			GenericResponse: api.NewClientError(),
			Payload:         ConfirmPayload{},
		}
	}
	var response UserPasswordResponse
	json.Unmarshal([]byte(body), &response)
	return response
}

func Link(context config.Context, data LinkRequest) LinkResponse {
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
		if context.Debug {
			fmt.Println("Client Request Error:")
			fmt.Println(errs)
		}
		return LinkResponse{
			GenericResponse: api.NewClientError(),
			Payload:         ConfirmPayload{},
		}
	}
	var response LinkResponse
	json.Unmarshal([]byte(body), &response)
	return response
}

func Unlink(context config.Context, data UnlinkRequest) UnlinkResponse {
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
		if context.Debug {
			fmt.Println("Client Request Error:")
			fmt.Println(errs)
		}
		return UnlinkResponse{
			GenericResponse: api.NewClientError(),
			Payload:         ConfirmPayload{},
		}
	}
	var response UnlinkResponse
	json.Unmarshal([]byte(body), &response)
	return response
}
