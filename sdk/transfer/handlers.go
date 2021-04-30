package transfer

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"mime"

	"github.com/jetrails/jrctl/sdk/utils"
	"github.com/jetrails/jrctl/sdk/version"
	"github.com/parnurzeal/gorequest"
)

func Send(context PublicApiContext, data SendRequest) (SendResponse, *ErrorResponse) {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	response, body, errors := request.
		Post(fmt.Sprintf("https://%s/%s", context.Endpoint, "transfer/upload")).
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Type("multipart").
		SendFile(data.FilePath).
		End()
	if errors != nil {
		return SendResponse{}, &ErrorResponse{Message: utils.CollectErrors(errors)[0]}
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
		return SendResponse{}, &errorResponse
	}
	var successResponse []SendResponse
	json.Unmarshal([]byte(body), &successResponse)
	return successResponse[0], nil
}

func Receive(context PublicApiContext, data ReceiveRequest) (ReceiveResponse, *ErrorResponse) {
	var request = gorequest.New()
	request.SetDebug(context.Debug)
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: context.Insecure})
	response, body, errors := request.
		Get(fmt.Sprintf("https://%s/%s", context.Endpoint, "transfer/download")).
		Set("User-Agent", fmt.Sprintf("jrctl/%s", version.VersionString)).
		Query(data).
		End()
	if errors != nil {
		return ReceiveResponse{}, &ErrorResponse{Message: utils.CollectErrors(errors)[0]}
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
		return ReceiveResponse{}, &errorResponse
	}
	disposition := response.Header.Get("Content-Disposition")
	filename := data.Identifier + "-" + data.Password
	if _, params, error := mime.ParseMediaType(disposition); error == nil {
		if params["filename"] != "" {
			filename = params["filename"]
		}
	}
	return ReceiveResponse{FileName: filename, Bytes: []byte(body)}, nil
}
