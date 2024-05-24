package server

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

func Version(context config.Context) VersionResponse {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	_, body, errs := request.
		Timeout(1*time.Second).
		Get(fmt.Sprintf("https://%s/version", context.Endpoint)).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Set("Authorization", "Bearer "+context.Token).
		Type("text").
		Send(`{}`).
		End()
	if len(errs) > 0 {
		if context.Debug {
			fmt.Println("Client Request Error:")
			fmt.Println(errs)
		}
		return VersionResponse{
			GenericResponse: api.NewClientError(),
			Payload:         "",
		}
	}
	var response VersionResponse
	json.Unmarshal([]byte(body), &response)
	return response
}

func TokenInfo(context config.Context) TokenResponse {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	_, body, errs := request.
		Timeout(1*time.Second).
		Get(fmt.Sprintf("https://%s/token", context.Endpoint)).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Set("Authorization", "Bearer "+context.Token).
		Type("text").
		Send(`{}`).
		End()
	if len(errs) > 0 {
		if context.Debug {
			fmt.Println("Client Request Error:")
			fmt.Println(errs)
		}
		return TokenResponse{
			GenericResponse: api.NewClientError(),
			Payload:         nil,
		}
	}
	var response TokenResponse
	json.Unmarshal([]byte(body), &response)
	return response
}
