package server

import (
	"fmt"
	"crypto/tls"
	"encoding/json"
	"github.com/jetrails/jrctl/sdk/version"
	"github.com/parnurzeal/gorequest"
)

func Version ( context Context ) VersionResponse {
	var request = gorequest.New ()
	request.SetDebug ( context.Debug )
	request.TLSClientConfig ( &tls.Config { InsecureSkipVerify: true })
	_, body, errors := request.
		Get ( fmt.Sprintf ("https://%s/version", context.Endpoint ) ).
		Set ( "Content-Type", "application/json" ).
		Set ( "User-Agent", fmt.Sprintf ( "jrctl/%s", version.VersionString ) ).
		Set ( "Authorization", context.Token ).
		Type ("text").
		Send (`{}`).
		End ()
	if len ( errors ) > 0 {
		return VersionResponse {
			Status: "Client Error",
			Code: 1,
			Messages: [] string { "Failed to connect to server." },
			Payload: "",
		}
	}
	var response VersionResponse
	json.Unmarshal ( [] byte ( body ), &response )
	return response
}

func ListServices ( context Context ) ListServicesResponse {
	var request = gorequest.New ()
	request.SetDebug ( context.Debug )
	request.TLSClientConfig ( &tls.Config { InsecureSkipVerify: true })
	_, body, errors := request.
		Get ( fmt.Sprintf ("https://%s/service", context.Endpoint ) ).
		Set ( "Content-Type", "application/json" ).
		Set ( "User-Agent", fmt.Sprintf ( "jrctl/%s", version.VersionString ) ).
		Set ( "Authorization", context.Token ).
		Type ("text").
		Send (`{}`).
		End ()
	if len ( errors ) > 0 {
		return ListServicesResponse {
			Status: "Client Error",
			Code: 1,
			Messages: [] string { "Failed to connect to server." },
			Payload: [] string {},
		}
	}
	var response ListServicesResponse
	json.Unmarshal ( [] byte ( body ), &response )
	return response
}