package transfer

import (
	"fmt"
	"encoding/json"
	"crypto/tls"
	"github.com/jetrails/jrctl/sdk/utils"
	"github.com/parnurzeal/gorequest"
)

func Send ( context PublicApiContext, data SendRequest ) ( SendResponse, * ErrorResponse ) {
	var request = gorequest.New ()
	request.SetDebug ( context.Debug )
	request.TLSClientConfig ( &tls.Config { InsecureSkipVerify: context.Insecure })
	response, body, errors := request.
		Post ( fmt.Sprintf ( "https://%s/%s", context.Endpoint, "transfer/upload" ) ).
		Type ("multipart").
		SendFile ( data.FilePath ).
		End ()
	if errors != nil {
		return SendResponse {}, &ErrorResponse { Message: utils.CollectErrors ( errors ) [ 0 ] }
	}
	if response.StatusCode != 200 {
		var errorResponse ErrorResponse
		json.Unmarshal ( [] byte ( body ), &errorResponse )
		if errorResponse.Code == 0 {
			errorResponse = ErrorResponse {
				Name: response.Status,
				Message: fmt.Sprintf ( "Endpoint: %s", response.Status ),
				Code: response.StatusCode,
				Type: response.Status,
				Data: nil,
			}
		}
		return SendResponse {}, &errorResponse
	}
	var successResponse [] SendResponse
	json.Unmarshal ( [] byte ( body ), &successResponse )
	return successResponse [0], nil
}

func Receive ( context PublicApiContext, data ReceiveRequest ) ( ReceiveResponse, * ErrorResponse ) {
	var request = gorequest.New ()
	request.SetDebug ( context.Debug )
	request.TLSClientConfig ( &tls.Config { InsecureSkipVerify: context.Insecure })
	response, body, errors := request.
		Get ( fmt.Sprintf ( "https://%s/%s", context.Endpoint, "transfer/download" ) ).
		Query ( data ).
		End ()
	if errors != nil {
		return ReceiveResponse {}, &ErrorResponse { Message: utils.CollectErrors ( errors ) [ 0 ] }
	}
	if response.StatusCode != 200 {
		var errorResponse ErrorResponse
		json.Unmarshal ( [] byte ( body ), &errorResponse )
		if errorResponse.Code == 0 {
			errorResponse = ErrorResponse {
				Name: response.Status,
				Message: fmt.Sprintf ( "Endpoint: %s", response.Status ),
				Code: response.StatusCode,
				Type: response.Status,
				Data: nil,
			}
		}
		return ReceiveResponse {}, &errorResponse
	}
	return ReceiveResponse { Bytes: [] byte ( body ) }, nil
}
