package firewall

import (
	"crypto/tls"
	"encoding/json"
	"fmt"

	"github.com/jetrails/jrctl/sdk/server"
	"github.com/jetrails/jrctl/sdk/version"
	"github.com/parnurzeal/gorequest"
)

func List(context server.Context) ListResponse {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	_, body, errors := request.
		Get(fmt.Sprintf("https://%s/whitelist", context.Endpoint)).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Set("Authorization", context.Token).
		Type("text").
		Send(`{}`).
		End()
	if len(errors) > 0 {
		return ListResponse{
			Status:   "Client Error",
			Code:     1,
			Messages: []string{"Failed to connect to server."},
			Payload:  nil,
		}
	}
	var response ListResponse
	json.Unmarshal([]byte(body), &response)
	return response
}

func Add(context server.Context, data AllowRequest) AllowResponse {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	_, body, errors := request.
		Put(fmt.Sprintf("https://%s/whitelist", context.Endpoint)).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Set("Authorization", context.Token).
		Send(data).
		End()
	if len(errors) > 0 {
		return AllowResponse{
			Status:   "Client Error",
			Code:     1,
			Messages: []string{"Failed to connect to server."},
			Payload:  data,
		}
	}
	var response AllowResponse
	json.Unmarshal([]byte(body), &response)
	return response
}

func Cloudflare(context server.Context, data CloudflareRequest) CloudflareResponse {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	_, body, errors := request.
		Put(fmt.Sprintf("https://%s/whitelist/cloudflare", context.Endpoint)).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Set("Authorization", context.Token).
		Send(data).
		End()
	if len(errors) > 0 {
		return CloudflareResponse{
			Status:   "Client Error",
			Code:     1,
			Messages: []string{"Failed to connect to server."},
			Payload: CloudflareEntries{
				Skipped: []string{},
				Allowed: []string{},
				Failed:  []CloudflareEntry{},
			},
		}
	}
	var response CloudflareResponse
	json.Unmarshal([]byte(body), &response)
	return response
}
