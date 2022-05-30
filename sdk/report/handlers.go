package report

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

func Audit(context config.Context) AuditResponse {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	_, body, errs := request.
		Timeout(5*time.Second).
		Get(fmt.Sprintf("https://%s/report/audit", context.Endpoint)).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Set("Authorization", "Bearer "+context.Token).
		Type("text").
		Send(`{}`).
		End()
	if len(errs) > 0 {
		return AuditResponse{
			GenericResponse: api.NewClientError(),
			Payload:         nil,
		}
	}
	var response AuditResponse
	json.Unmarshal([]byte(body), &response)
	return response
}
