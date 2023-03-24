package alternative

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
		Timeout(10*time.Second).
		Get(fmt.Sprintf("https://%s/alternative", context.Endpoint)).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Set("Authorization", "Bearer "+context.Token).
		Type("text").
		Send("{}").
		End()
	if len(errs) > 0 {
		return ListResponse{
			GenericResponse: api.NewClientError(),
			Payload:         []Entry{},
		}
	}
	var response ListResponse
	json.Unmarshal([]byte(body), &response)
	return response
}

func Switch(context config.Context, data SwitchRequest) SwitchResponse {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	_, body, errs := request.
		Timeout(10*time.Second).
		Put(fmt.Sprintf("https://%s/alternative", context.Endpoint)).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Set("Authorization", "Bearer "+context.Token).
		Send(data).
		End()
	if len(errs) > 0 {
		return SwitchResponse{
			GenericResponse: api.NewClientError(),
			Payload:         nil,
		}
	}
	var response SwitchResponse
	json.Unmarshal([]byte(body), &response)
	return response
}
