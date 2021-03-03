package firewall

import (
	"fmt"
	"crypto/tls"
	"encoding/json"
	"github.com/jetrails/jrctl/sdk/version"
	"github.com/jetrails/jrctl/sdk/daemon"
	"github.com/parnurzeal/gorequest"
)

func List ( context daemon.Context ) ListResponse {
	var request = gorequest.New ()
	request.SetDebug ( context.Debug )
	request.TLSClientConfig ( &tls.Config { InsecureSkipVerify: true })
	_, body, errors := request.
		Get ( fmt.Sprintf ("https://%s/whitelist", context.Endpoint ) ).
		Set ( "Content-Type", "application/json" ).
		Set ( "User-Agent", fmt.Sprintf ( "jrctl/%s", version.VersionString ) ).
		Set ( "Authorization", context.Token ).
		Type ("text").
		Send (`{}`).
		End ()
	if len ( errors ) > 0 {
		return ListResponse {
			Status: "Client Side",
			Code: 1,
			Messages: [] string { "Failed to connect to daemon." },
			Payload: nil,
		}
	}
	var response ListResponse
	json.Unmarshal ( [] byte ( body ), &response )
	return response
}

func Add ( context daemon.Context, data AllowRequest ) AllowResponse {
	var request = gorequest.New ()
	request.SetDebug ( context.Debug )
	request.TLSClientConfig ( &tls.Config { InsecureSkipVerify: true })
	_, body, errors := request.
		Put ( fmt.Sprintf ("https://%s/whitelist", context.Endpoint ) ).
		Set ( "Content-Type", "application/json" ).
		Set ( "User-Agent", fmt.Sprintf ( "jrctl/%s", version.VersionString ) ).
		Set ( "Authorization", context.Token ).
		Send ( data ).
		End ()
	if len ( errors ) > 0 {
		return AllowResponse {
			Status: "Client Side",
			Code: 1,
			Messages: [] string { "Failed to connect to daemon." },
			Payload: data,
		}
	}
	var response AllowResponse
	json.Unmarshal ( [] byte ( body ), &response )
	return response
}
