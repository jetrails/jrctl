package service

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

func ListServices(context config.Context) ListServicesResponse {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	_, body, errs := request.
		Timeout(10*time.Second).
		Get(fmt.Sprintf("https://%s/service", context.Endpoint)).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Set("Authorization", "Bearer "+context.Token).
		Type("text").
		Send(`{}`).
		End()
	if len(errs) > 0 {
		if context.Debug {
			fmt.Println("Error in ListServices:")
			fmt.Println(errs)
		}
		return ListServicesResponse{
			GenericResponse: api.NewClientError(),
			Payload:         map[string]ServiceProperties{},
		}
	}
	var response ListServicesResponse
	json.Unmarshal([]byte(body), &response)
	return response
}

func Restart(context config.Context, data RestartRequest) RestartResponse {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	_, body, errs := request.
		Timeout(30*time.Second).
		Put(fmt.Sprintf("https://%s/service/restart", context.Endpoint)).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Set("Authorization", "Bearer "+context.Token).
		Send(data).
		End()
	if len(errs) > 0 {
		return RestartResponse{
			GenericResponse: api.NewClientError(),
			Payload:         data,
		}
	}
	var response RestartResponse
	json.Unmarshal([]byte(body), &response)
	return response
}

func Reload(context config.Context, data ReloadRequest) ReloadResponse {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	_, body, errs := request.
		Timeout(30*time.Second).
		Put(fmt.Sprintf("https://%s/service/reload", context.Endpoint)).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Set("Authorization", "Bearer "+context.Token).
		Send(data).
		End()
	if len(errs) > 0 {
		return ReloadResponse{
			GenericResponse: api.NewClientError(),
			Payload:         data,
		}
	}
	var response ReloadResponse
	json.Unmarshal([]byte(body), &response)
	return response
}

func Enable(context config.Context, data EnableRequest) EnableResponse {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	_, body, errs := request.
		Timeout(30*time.Second).
		Put(fmt.Sprintf("https://%s/service/enable", context.Endpoint)).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Set("Authorization", "Bearer "+context.Token).
		Send(data).
		End()
	if len(errs) > 0 {
		return EnableResponse{
			GenericResponse: api.NewClientError(),
			Payload:         data,
		}
	}
	var response EnableResponse
	json.Unmarshal([]byte(body), &response)
	return response
}

func Disable(context config.Context, data DisableRequest) DisableResponse {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	_, body, errs := request.
		Timeout(30*time.Second).
		Put(fmt.Sprintf("https://%s/service/disable", context.Endpoint)).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Set("Authorization", "Bearer "+context.Token).
		Send(data).
		End()
	if len(errs) > 0 {
		return DisableResponse{
			GenericResponse: api.NewClientError(),
			Payload:         data,
		}
	}
	var response DisableResponse
	json.Unmarshal([]byte(body), &response)
	return response
}
