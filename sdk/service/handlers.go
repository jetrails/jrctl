package service

import (
	"fmt"
	"crypto/tls"
	"encoding/json"
	"github.com/jetrails/jrctl/sdk/version"
	"github.com/jetrails/jrctl/sdk/daemon"
	"github.com/parnurzeal/gorequest"
)

func Restart ( context daemon.Context, data RestartRequest ) RestartResponse {
	var request = gorequest.New ()
	request.SetDebug ( context.Debug )
	request.TLSClientConfig ( &tls.Config { InsecureSkipVerify: true })
	_, body, errors := request.
		Post ( fmt.Sprintf ("https://%s/service/restart", context.Endpoint ) ).
		Set ( "Content-Type", "application/json" ).
		Set ( "User-Agent", fmt.Sprintf ( "jrctl/%s", version.VersionString ) ).
		Set ( "Authorization", context.Token ).
		Send ( data ).
		End ()
	if len ( errors ) > 0 {
		return RestartResponse {
			Status: "Client Error",
			Code: 1,
			Messages: [] string { "Failed to connect to daemon." },
			Payload: data,
		}
	}
	var response RestartResponse
	json.Unmarshal ( [] byte ( body ), &response )
	return response
}
