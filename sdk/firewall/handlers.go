package firewall

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
		Timeout(5*time.Second).
		Get(fmt.Sprintf("https://%s/firewall", context.Endpoint)).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Set("Authorization", "Bearer "+context.Token).
		Type("text").
		Send(`{}`).
		End()
	if len(errs) > 0 {
		return ListResponse{
			Status:   "Client Error",
			Code:     1,
			Messages: []string{"Failed to connect to server."},
			Payload:  nil,
		}
	}
	var response ListResponse
	json.Unmarshal([]byte(body), &response)
	if len(response.Messages) == 0 {
		response.Messages = append(response.Messages, "Endpoint: "+response.Status)
	}
	return response
}

func Allow(context server.Context, data AllowRequest) AllowResponse {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	_, body, errs := request.
		Timeout(3*time.Second).
		Put(fmt.Sprintf("https://%s/firewall/allow", context.Endpoint)).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Set("Authorization", "Bearer "+context.Token).
		Send(data).
		End()
	if len(errs) > 0 {
		return AllowResponse{
			Status:   "Client Error",
			Code:     1,
			Messages: []string{"Failed to connect to server."},
			Payload:  data,
		}
	}
	var response AllowResponse
	json.Unmarshal([]byte(body), &response)
	if len(response.Messages) == 0 {
		response.Messages = append(response.Messages, "Endpoint: "+response.Status)
	}
	return response
}

func UnAllow(context server.Context, data UnAllowRequest) UnAllowResponse {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	_, body, errs := request.
		Timeout(3*time.Second).
		Delete(fmt.Sprintf("https://%s/firewall/allow", context.Endpoint)).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Set("Authorization", "Bearer "+context.Token).
		Send(data).
		End()
	if len(errs) > 0 {
		return UnAllowResponse{
			Status:   "Client Error",
			Code:     1,
			Messages: []string{"Failed to connect to server."},
			Payload:  data,
		}
	}
	var response UnAllowResponse
	json.Unmarshal([]byte(body), &response)
	if len(response.Messages) == 0 {
		response.Messages = append(response.Messages, "Endpoint: "+response.Status)
	}
	return response
}

func Deny(context server.Context, data DenyRequest) DenyResponse {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	_, body, errs := request.
		Timeout(3*time.Second).
		Put(fmt.Sprintf("https://%s/firewall/deny", context.Endpoint)).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Set("Authorization", "Bearer "+context.Token).
		Send(data).
		End()
	if len(errs) > 0 {
		return DenyResponse{
			Status:   "Client Error",
			Code:     1,
			Messages: []string{"Failed to connect to server."},
			Payload:  data,
		}
	}
	var response DenyResponse
	json.Unmarshal([]byte(body), &response)
	if len(response.Messages) == 0 {
		response.Messages = append(response.Messages, "Endpoint: "+response.Status)
	}
	return response
}

func UnDeny(context server.Context, data UnDenyRequest) UnDenyResponse {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	_, body, errs := request.
		Timeout(3*time.Second).
		Delete(fmt.Sprintf("https://%s/firewall/deny", context.Endpoint)).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Set("Authorization", "Bearer "+context.Token).
		Send(data).
		End()
	if len(errs) > 0 {
		return UnDenyResponse{
			Status:   "Client Error",
			Code:     1,
			Messages: []string{"Failed to connect to server."},
			Payload:  data,
		}
	}
	var response UnDenyResponse
	json.Unmarshal([]byte(body), &response)
	if len(response.Messages) == 0 {
		response.Messages = append(response.Messages, "Endpoint: "+response.Status)
	}
	return response
}

func AllowCloudflare(context server.Context) AllowCloudflareResponse {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	_, body, errs := request.
		Timeout(15*time.Second).
		Put(fmt.Sprintf("https://%s/firewall/allow/cloudflare", context.Endpoint)).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Set("Authorization", "Bearer "+context.Token).
		Type("text").
		Send(`{}`).
		End()
	if len(errs) > 0 {
		return AllowCloudflareResponse{
			Status:   "Client Error",
			Code:     1,
			Messages: []string{"Failed to connect to server."},
			Payload: CloudflareEntries{
				Skipped:   []string{},
				Succeeded: []string{},
				Failed:    []CloudflareEntry{},
			},
		}
	}
	var response AllowCloudflareResponse
	json.Unmarshal([]byte(body), &response)
	if len(response.Messages) == 0 {
		response.Messages = append(response.Messages, "Endpoint: "+response.Status)
	}
	return response
}

func UnAllowCloudflare(context server.Context) UnAllowCloudflareResponse {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	_, body, errs := request.
		Timeout(15*time.Second).
		Delete(fmt.Sprintf("https://%s/firewall/allow/cloudflare", context.Endpoint)).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Set("Authorization", "Bearer "+context.Token).
		Type("text").
		Send(`{}`).
		End()
	if len(errs) > 0 {
		return UnAllowCloudflareResponse{
			Status:   "Client Error",
			Code:     1,
			Messages: []string{"Failed to connect to server."},
			Payload: CloudflareEntries{
				Skipped:   []string{},
				Succeeded: []string{},
				Failed:    []CloudflareEntry{},
			},
		}
	}
	var response UnAllowCloudflareResponse
	json.Unmarshal([]byte(body), &response)
	if len(response.Messages) == 0 {
		response.Messages = append(response.Messages, "Endpoint: "+response.Status)
	}
	return response
}
