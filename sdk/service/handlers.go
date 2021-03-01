package service

import (
	"fmt"
	"crypto/tls"
	"encoding/json"
	"github.com/jetrails/jrctl/sdk/utils"
	"github.com/jetrails/jrctl/sdk/version"
	"github.com/parnurzeal/gorequest"
)

func Restart ( context DaemonContext, data RestartRequest ) RestartResponse {
	var request = gorequest.New ()
	request.SetDebug ( context.Debug )
	request.TLSClientConfig ( &tls.Config { InsecureSkipVerify: true })
	_, body, errors := request.
		Post ( fmt.Sprintf ("https://%s/service/restart", context.Endpoint ) ).
		Set ( "Content-Type", "application/json" ).
		Set ( "User-Agent", fmt.Sprintf ( "jrctl/%s", version.VersionString ) ).
		Set ( "Authorization", context.Auth ).
		Send ( data ).
		End ()
	if len ( errors ) > 0 {
		return RestartResponse {
			Status: "Client Side",
			Code: 1,
			Messages: utils.CollectErrors ( errors ),
			Payload: data,
		}
	}
	var response RestartResponse
	json.Unmarshal ( [] byte ( body ), &response )
	return response
}
