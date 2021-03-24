package transfer

import (
	"fmt"
	"encoding/json"
	"github.com/jetrails/jrctl/sdk/utils"
	"github.com/parnurzeal/gorequest"
)

func Send ( context PublicApiContext, data SendRequest ) ( SendResponse, ErrorResponse ) {
	var request = gorequest.New ()
	request.SetDebug ( context.Debug )
	response, body, errors := request.
		Post ( fmt.Sprintf ( "https://%s/%s", context.Endpoint, "transfer/upload" ) ).
		Type ("multipart").
		SendFile ( data.FilePath ).
		End ()
	if errors != nil {
		return SendResponse {}, ErrorResponse { Message: utils.CollectErrors ( errors ) [ 0 ] }
	}
	if response.StatusCode != 200 {
		var errorResponse ErrorResponse
		json.Unmarshal ( [] byte ( body ), &errorResponse )
		return SendResponse {}, errorResponse
	}
	var successResponse [] SendResponse
	json.Unmarshal ( [] byte ( body ), &successResponse )
	return successResponse [0], ErrorResponse {}
}

func Receive ( context PublicApiContext, data ReceiveRequest ) ( ReceiveResponse, ErrorResponse ) {
	var request = gorequest.New ()
	request.SetDebug ( context.Debug )
	response, body, errors := request.
		Get ( fmt.Sprintf ( "https://%s/%s", context.Endpoint, "transfer/download" ) ).
		Query ( data ).
		End ()
	if errors != nil {
		return ReceiveResponse {}, ErrorResponse { Message: utils.CollectErrors ( errors ) [ 0 ] }
	}
	if response.StatusCode != 200 {
		var errorResponse ErrorResponse
		json.Unmarshal ( [] byte ( body ), &errorResponse )
		return ReceiveResponse {}, errorResponse
	}
	return ReceiveResponse { Bytes: [] byte ( body ) }, ErrorResponse {}
}
