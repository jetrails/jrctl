package secret

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jetrails/jrctl/sdk/utils"
	"github.com/jetrails/jrctl/sdk/version"
	"github.com/parnurzeal/gorequest"
)

func SecretCreate(context PublicApiContext, data SecretCreateRequest) (SecretCreateResponse, *ErrorResponse) {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	response, body, errs := request.
		Timeout(10*time.Second).
		Post(fmt.Sprintf("https://%s/%s", context.Endpoint, "secret")).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Send(data).
		End()
	if errs != nil {
		return SecretCreateResponse{}, &ErrorResponse{Message: utils.CollectErrors(errs)[0]}
	}
	if response != nil && response.StatusCode != 200 {
		var errorResponse ErrorResponse
		json.Unmarshal([]byte(body), &errorResponse)
		if errorResponse.Code == 0 {
			errorResponse = ErrorResponse{
				Name:    response.Status,
				Message: fmt.Sprintf("Endpoint: %s", response.Status),
				Code:    response.StatusCode,
				Type:    response.Status,
				Data:    nil,
			}
		}
		return SecretCreateResponse{}, &errorResponse
	}
	var successResponse SecretCreateResponse
	json.Unmarshal([]byte(body), &successResponse)
	return successResponse, nil
}

func SecretDelete(context PublicApiContext, data SecretDeleteRequest) (SecretDeleteResponse, *ErrorResponse) {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	response, body, errs := request.
		Timeout(10*time.Second).
		Delete(fmt.Sprintf("https://%s/%s", context.Endpoint, "secret")).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Send(data).
		End()
	if errs != nil {
		return SecretDeleteResponse{}, &ErrorResponse{Message: utils.CollectErrors(errs)[0]}
	}
	if response != nil && response.StatusCode != 200 {
		var errorResponse ErrorResponse
		json.Unmarshal([]byte(body), &errorResponse)
		if errorResponse.Code == 0 {
			errorResponse = ErrorResponse{
				Name:    response.Status,
				Message: fmt.Sprintf("Endpoint: %s", response.Status),
				Code:    response.StatusCode,
				Type:    response.Status,
				Data:    nil,
			}
		}
		return SecretDeleteResponse{}, &errorResponse
	}
	var successResponse SecretDeleteResponse
	json.Unmarshal([]byte(body), &successResponse)
	return successResponse, nil
}

func SecretRead(context PublicApiContext, data SecretReadRequest) (SecretReadResponse, *ErrorResponse) {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	response, body, errs := request.
		Timeout(10*time.Second).
		Get(fmt.Sprintf("https://%s/%s", context.Endpoint, "secret")).
		Set("Content-Type", "application/json").
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Query(data).
		End()
	if errs != nil {
		return SecretReadResponse{}, &ErrorResponse{Message: utils.CollectErrors(errs)[0]}
	}
	if response != nil && response.StatusCode != 200 {
		var errorResponse ErrorResponse
		json.Unmarshal([]byte(body), &errorResponse)
		if errorResponse.Code == 0 {
			errorResponse = ErrorResponse{
				Name:    response.Status,
				Message: fmt.Sprintf("Endpoint: %s", response.Status),
				Code:    response.StatusCode,
				Type:    response.Status,
				Data:    nil,
			}
		}
		return SecretReadResponse{}, &errorResponse
	}
	return SecretReadResponse{Data: string(body)}, nil
}
