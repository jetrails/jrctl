package secret

import "encoding/json"
import "github.com/jetrails/jrctl/sdk/utils"

type SecretCreateRequest struct {
	Data string             `json:"data"`
	Password string         `json:"password,omitempty"`
	TTL int                 `json:"ttl,omitempty"`
}

type SecretReadRequest struct {
	Identifier string       `json:"id"`
	Password string         `json:"password"`
}

type SecretDeleteRequest struct {
	Identifier string       `json:"id"`
}

type SecretCreateResponse struct {
	Identifier string       `json:"id"`
	Password string         `json:"password"`
	TTL int                 `json:"ttl"`
}

type SecretReadResponse struct {
	Data string             `json:"data"`
}

type SecretDeleteResponse struct {
	Identifier string       `json:"id"`
}

func SecretCreate ( request SecretCreateRequest ) ( SecretCreateResponse, utils.ErrorResponse ) {
	response, body, error := utils.Request ().
		Post ( utils.GetPublicEndpoint ("secret") ).
		Send ( request ).
		End ()
	if error != nil {
		return SecretCreateResponse {}, utils.InternalError ( error )
	}
	if response.StatusCode != 200 {
		var errorResponse utils.ErrorResponse
		json.Unmarshal ( [] byte ( body ), &errorResponse )
		return SecretCreateResponse {}, errorResponse
	}
	var successResponse SecretCreateResponse
	json.Unmarshal ( [] byte ( body ), &successResponse )
	return successResponse, utils.ErrorResponse {}
}

func SecretDelete ( request SecretDeleteRequest ) ( SecretDeleteResponse, utils.ErrorResponse ) {
	response, body, error := utils.Request ().
		Delete ( utils.GetPublicEndpoint ("secret") ).
		Send ( request ).
		End ()
	if error != nil {
		return SecretDeleteResponse {}, utils.InternalError ( error )
	}
	if response.StatusCode != 200 {
		var errorResponse utils.ErrorResponse
		json.Unmarshal ( [] byte ( body ), &errorResponse )
		return SecretDeleteResponse {}, errorResponse
	}
	var successResponse SecretDeleteResponse
	json.Unmarshal ( [] byte ( body ), &successResponse )
	return successResponse, utils.ErrorResponse {}
}

func SecretRead ( request SecretReadRequest ) ( SecretReadResponse, utils.ErrorResponse ) {
	response, body, error := utils.Request ().
		Get ( utils.GetPublicEndpoint ("secret") ).
		Query ( request ).
		End ()
	if error != nil {
		return SecretReadResponse {}, utils.InternalError ( error )
	}
	if response.StatusCode != 200 {
		var errorResponse utils.ErrorResponse
		json.Unmarshal ( [] byte ( body ), &errorResponse )
		return SecretReadResponse {}, errorResponse
	}
	return SecretReadResponse { Data: string ( body ) }, utils.ErrorResponse {}
}
