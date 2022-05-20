package database

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jetrails/jrctl/sdk/server"
	"github.com/jetrails/jrctl/sdk/version"
	"github.com/parnurzeal/gorequest"
)

func List(context server.Context) ListResponse {
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
		return ListResponse{
			Status:   "Client Error",
			Code:     1,
			Messages: []string{"Failed to connect to server."},
			Payload:  []Database{},
		}
	}
	var response ListResponse
	json.Unmarshal([]byte(body), &response)
	if len(response.Messages) == 0 {
		response.Messages = append(response.Messages, "Endpoint: "+response.Status)
	}
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
			Status:   "Client Error",
			Code:     1,
			Messages: []string{"Failed to connect to server."},
			Payload:  "",
		}
	}
	var response CreateResponse
	json.Unmarshal([]byte(body), &response)
	if len(response.Messages) == 0 {
		response.Messages = append(response.Messages, "Endpoint: "+response.Status)
	}
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
			Status:   "Client Error",
			Code:     1,
			Messages: []string{"Failed to connect to server."},
			Payload:  "",
		}
	}
	var response DeleteResponse
	json.Unmarshal([]byte(body), &response)
	if len(response.Messages) == 0 {
		response.Messages = append(response.Messages, "Endpoint: "+response.Status)
	}
	return response
}

func UserAdd(context server.Context, data UserAddRequest) UserAddResponse {
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
		return UserAddResponse{
			Status:   "Client Error",
			Code:     1,
			Messages: []string{"Failed to connect to server."},
			Payload:  "",
		}
	}
	var response UserAddResponse
	json.Unmarshal([]byte(body), &response)
	if len(response.Messages) == 0 {
		response.Messages = append(response.Messages, "Endpoint: "+response.Status)
	}
	return response
}

func UserRemove(context server.Context, data UserRemoveRequest) UserRemoveResponse {
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
		return UserRemoveResponse{
			Status:   "Client Error",
			Code:     1,
			Messages: []string{"Failed to connect to server."},
			Payload:  "",
		}
	}
	var response UserRemoveResponse
	json.Unmarshal([]byte(body), &response)
	if len(response.Messages) == 0 {
		response.Messages = append(response.Messages, "Endpoint: "+response.Status)
	}
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
			Status:   "Client Error",
			Code:     1,
			Messages: []string{"Failed to connect to server."},
			Payload:  "",
		}
	}
	var response UserPasswordResponse
	json.Unmarshal([]byte(body), &response)
	if len(response.Messages) == 0 {
		response.Messages = append(response.Messages, "Endpoint: "+response.Status)
	}
	return response
}
