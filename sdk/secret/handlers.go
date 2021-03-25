package secret

import (
	"fmt"
	"encoding/json"
	"crypto/tls"
	"github.com/jetrails/jrctl/sdk/utils"
	"github.com/parnurzeal/gorequest"
)

func SecretCreate ( context PublicApiContext, data SecretCreateRequest ) ( SecretCreateResponse, * ErrorResponse ) {
	var request = gorequest.New ()
	request.SetDebug ( context.Debug )
	request.TLSClientConfig ( &tls.Config { InsecureSkipVerify: context.Insecure })
	response, body, errors := request.
		Post ( fmt.Sprintf ( "https://%s/%s", context.Endpoint, "secret" ) ).
		Send ( data ).
		End ()
	if errors != nil {
		return SecretCreateResponse {}, &ErrorResponse { Message: utils.CollectErrors ( errors ) [ 0 ] }
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
		return SecretCreateResponse {}, &errorResponse
	}
	var successResponse SecretCreateResponse
	json.Unmarshal ( [] byte ( body ), &successResponse )
	return successResponse, nil
}

func SecretDelete ( context PublicApiContext, data SecretDeleteRequest ) ( SecretDeleteResponse, * ErrorResponse ) {
	var request = gorequest.New ()
	request.SetDebug ( context.Debug )
	request.TLSClientConfig ( &tls.Config { InsecureSkipVerify: context.Insecure })
	response, body, errors := request.
		Delete ( fmt.Sprintf ( "https://%s/%s", context.Endpoint, "secret" ) ).
		Send ( data ).
		End ()
	if errors != nil {
		return SecretDeleteResponse {}, &ErrorResponse { Message: utils.CollectErrors ( errors ) [ 0 ] }
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
		return SecretDeleteResponse {}, &errorResponse
	}
	var successResponse SecretDeleteResponse
	json.Unmarshal ( [] byte ( body ), &successResponse )
	return successResponse, nil
}

func SecretRead ( context PublicApiContext, data SecretReadRequest ) ( SecretReadResponse, * ErrorResponse ) {
	var request = gorequest.New ()
	request.SetDebug ( context.Debug )
	request.TLSClientConfig ( &tls.Config { InsecureSkipVerify: context.Insecure })
	response, body, errors := request.
		Get ( fmt.Sprintf ( "https://%s/%s", context.Endpoint, "secret" ) ).
		Query ( data ).
		End ()
	if errors != nil {
		return SecretReadResponse {}, &ErrorResponse { Message: utils.CollectErrors ( errors ) [ 0 ] }
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
		return SecretReadResponse {}, &errorResponse
	}
	return SecretReadResponse { Data: string ( body ) }, nil
}
