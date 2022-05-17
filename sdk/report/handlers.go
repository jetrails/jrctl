package report

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jetrails/jrctl/sdk/server"
	"github.com/jetrails/jrctl/sdk/version"
	"github.com/parnurzeal/gorequest"
)

func Audit(context server.Context) AuditResponse {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	_, body, errors := request.
		Timeout(5*time.Second).
		Get(fmt.Sprintf("https://%s/report/audit", context.Endpoint)).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Set("Authorization", "Bearer "+context.Token).
		Type("text").
		Send(`{}`).
		End()
	if len(errors) > 0 {
		return AuditResponse{
			Status:   "Client Error",
			Code:     1,
			Messages: []string{"Failed to connect to server."},
			Payload:  nil,
		}
	}
	var response AuditResponse
	json.Unmarshal([]byte(body), &response)
	if len(response.Messages) == 0 {
		response.Messages = append(response.Messages, "Endpoint: "+response.Status)
	}
	return response
}