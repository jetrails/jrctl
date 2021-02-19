package firewall

import (
	"fmt"
	"crypto/tls"
	"encoding/json"
	"github.com/jetrails/jrctl/sdk/utils"
	"github.com/jetrails/jrctl/sdk"
	"github.com/parnurzeal/gorequest"
)

func List ( context DaemonContext ) ListResponse {
	var request = gorequest.New ()
	request.SetDebug ( context.Debug )
	request.TLSClientConfig ( &tls.Config { InsecureSkipVerify: true })
	_, body, errors := request.
		Get ( fmt.Sprintf ("https://%s/whitelist", context.Endpoint ) ).
		Set ( "Content-Type", "application/json" ).
		Set ( "User-Agent", fmt.Sprintf ( "jrctl/%s", sdk.Version ) ).
		Set ( "Authorization", context.Auth ).
		Type ("text").
		Send (`{}`).
		End ()
	if len ( errors ) > 0 {
		return ListResponse {
			Status: "Client Side",
			Code: 1,
			Messages: utils.CollectErrors ( errors ),
			Payload: nil,
		}
	}
	var response ListResponse
	json.Unmarshal ( [] byte ( body ), &response )
	return response
}

func Add ( context DaemonContext, data AllowRequest ) AllowResponse {
	var request = gorequest.New ()
	request.SetDebug ( context.Debug )
	request.TLSClientConfig ( &tls.Config { InsecureSkipVerify: true })
	_, body, errors := request.
		Put ( fmt.Sprintf ("https://%s/whitelist", context.Endpoint ) ).
		Set ( "Content-Type", "application/json" ).
		Set ( "User-Agent", fmt.Sprintf ( "jrctl/%s", sdk.Version ) ).
		Set ( "Authorization", context.Auth ).
		Send ( data ).
		End ()
	if len ( errors ) > 0 {
		return AllowResponse {
			Status: "Client Side",
			Code: 1,
			Messages: utils.CollectErrors ( errors ),
			Payload: data,
		}
	}
	var response AllowResponse
	json.Unmarshal ( [] byte ( body ), &response )
	return response
}
