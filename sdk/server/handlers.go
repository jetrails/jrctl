package server

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jetrails/jrctl/sdk/version"
	"github.com/parnurzeal/gorequest"
)

func Version(context Context) VersionResponse {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	_, body, errors := request.
		Timeout(1*time.Second).
		Get(fmt.Sprintf("https://%s/version", context.Endpoint)).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Set("Authorization", "Bearer "+context.Token).
		Type("text").
		Send(`{}`).
		End()
	if len(errors) > 0 {
		return VersionResponse{
			Status:   "Client Error",
			Code:     1,
			Messages: []string{"Failed to connect to server."},
			Payload:  "",
		}
	}
	var response VersionResponse
	json.Unmarshal([]byte(body), &response)
	return response
}

func TokenInfo(context Context) TokenResponse {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	_, body, errors := request.
		Timeout(1*time.Second).
		Get(fmt.Sprintf("https://%s/token", context.Endpoint)).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Set("Authorization", "Bearer "+context.Token).
		Type("text").
		Send(`{}`).
		End()
	if len(errors) > 0 {
		return TokenResponse{
			Status:   "Client Error",
			Code:     1,
			Messages: []string{"Failed to connect to server."},
			Payload:  nil,
		}
	}
	var response TokenResponse
	json.Unmarshal([]byte(body), &response)
	return response
}

func ListServices(context Context) ListServicesResponse {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	_, body, errors := request.
		Timeout(1*time.Second).
		Get(fmt.Sprintf("https://%s/service", context.Endpoint)).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Set("Authorization", "Bearer "+context.Token).
		Type("text").
		Send(`{}`).
		End()
	if len(errors) > 0 {
		return ListServicesResponse{
			Status:   "Client Error",
			Code:     1,
			Messages: []string{"Failed to connect to server."},
			Payload:  map[string]ServiceProperties{},
		}
	}
	var response ListServicesResponse
	json.Unmarshal([]byte(body), &response)
	return response
}

func Restart(context Context, data RestartRequest) RestartResponse {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	_, body, errors := request.
		Timeout(30*time.Second).
		Put(fmt.Sprintf("https://%s/service/restart", context.Endpoint)).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Set("Authorization", "Bearer "+context.Token).
		Send(data).
		End()
	if len(errors) > 0 {
		return RestartResponse{
			Status:   "Client Error",
			Code:     1,
			Messages: []string{"Failed to connect to server."},
			Payload:  data,
		}
	}
	var response RestartResponse
	json.Unmarshal([]byte(body), &response)
	return response
}

func Reload(context Context, data ReloadRequest) ReloadResponse {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	_, body, errors := request.
		Timeout(30*time.Second).
		Put(fmt.Sprintf("https://%s/service/reload", context.Endpoint)).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Set("Authorization", "Bearer "+context.Token).
		Send(data).
		End()
	if len(errors) > 0 {
		return ReloadResponse{
			Status:   "Client Error",
			Code:     1,
			Messages: []string{"Failed to connect to server."},
			Payload:  data,
		}
	}
	var response ReloadResponse
	json.Unmarshal([]byte(body), &response)
	return response
}

func Enable(context Context, data EnableRequest) EnableResponse {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	_, body, errors := request.
		Timeout(30*time.Second).
		Put(fmt.Sprintf("https://%s/service/enable", context.Endpoint)).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Set("Authorization", "Bearer "+context.Token).
		Send(data).
		End()
	if len(errors) > 0 {
		return EnableResponse{
			Status:   "Client Error",
			Code:     1,
			Messages: []string{"Failed to connect to server."},
			Payload:  data,
		}
	}
	var response EnableResponse
	json.Unmarshal([]byte(body), &response)
	return response
}

func Disable(context Context, data DisableRequest) DisableResponse {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	_, body, errors := request.
		Timeout(30*time.Second).
		Put(fmt.Sprintf("https://%s/service/disable", context.Endpoint)).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Set("Authorization", "Bearer "+context.Token).
		Send(data).
		End()
	if len(errors) > 0 {
		return DisableResponse{
			Status:   "Client Error",
			Code:     1,
			Messages: []string{"Failed to connect to server."},
			Payload:  data,
		}
	}
	var response DisableResponse
	json.Unmarshal([]byte(body), &response)
	return response
}
