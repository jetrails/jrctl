package firewall

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

func List(context config.Context) ListResponse {
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
		if context.Debug {
			fmt.Println("Client Request Error:")
			fmt.Println(errs)
		}
		return ListResponse{
			GenericResponse: api.NewClientError(),
			Payload:         nil,
		}
	}
	var response ListResponse
	json.Unmarshal([]byte(body), &response)
	return response
}

func Allow(context config.Context, data AllowRequest) AllowResponse {
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
		if context.Debug {
			fmt.Println("Client Request Error:")
			fmt.Println(errs)
		}
		return AllowResponse{
			GenericResponse: api.NewClientError(),
			Payload:         data,
		}
	}
	var response AllowResponse
	json.Unmarshal([]byte(body), &response)
	return response
}

func UnAllow(context config.Context, data UnAllowRequest) UnAllowResponse {
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
		if context.Debug {
			fmt.Println("Client Request Error:")
			fmt.Println(errs)
		}
		return UnAllowResponse{
			GenericResponse: api.NewClientError(),
			Payload:         data,
		}
	}
	var response UnAllowResponse
	json.Unmarshal([]byte(body), &response)
	return response
}

func Deny(context config.Context, data DenyRequest) DenyResponse {
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
		if context.Debug {
			fmt.Println("Client Request Error:")
			fmt.Println(errs)
		}
		return DenyResponse{
			GenericResponse: api.NewClientError(),
			Payload:         data,
		}
	}
	var response DenyResponse
	json.Unmarshal([]byte(body), &response)
	return response
}

func UnDeny(context config.Context, data UnDenyRequest) UnDenyResponse {
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
		if context.Debug {
			fmt.Println("Client Request Error:")
			fmt.Println(errs)
		}
		return UnDenyResponse{
			GenericResponse: api.NewClientError(),
			Payload:         data,
		}
	}
	var response UnDenyResponse
	json.Unmarshal([]byte(body), &response)
	return response
}

func AllowCloudflare(context config.Context) AllowCloudflareResponse {
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
		if context.Debug {
			fmt.Println("Client Request Error:")
			fmt.Println(errs)
		}
		return AllowCloudflareResponse{
			GenericResponse: api.NewClientError(),
			Payload: CloudflareEntries{
				Skipped:   []string{},
				Succeeded: []string{},
				Failed:    []CloudflareEntry{},
			},
		}
	}
	var response AllowCloudflareResponse
	json.Unmarshal([]byte(body), &response)
	return response
}

func UnAllowCloudflare(context config.Context) UnAllowCloudflareResponse {
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
		if context.Debug {
			fmt.Println("Client Request Error:")
			fmt.Println(errs)
		}
		return UnAllowCloudflareResponse{
			GenericResponse: api.NewClientError(),
			Payload: CloudflareEntries{
				Skipped:   []string{},
				Succeeded: []string{},
				Failed:    []CloudflareEntry{},
			},
		}
	}
	var response UnAllowCloudflareResponse
	json.Unmarshal([]byte(body), &response)
	return response
}
