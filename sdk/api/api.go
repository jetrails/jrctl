package api

type Response interface {
	GetGeneric() *GenericResponse
}

type GenericResponse struct {
	Status   string            `json:"status"`
	Code     int               `json:"code"`
	Messages []string          `json:"messages"`
	Metadata map[string]string `json:"metadata"`
	Payload  interface{}       `json:"payload"`
}

func (res *GenericResponse) GetGeneric() *GenericResponse {
	return res
}

func (res *GenericResponse) IsOkay() bool {
	return res.Code == 200
}

func (res *GenericResponse) GetFirstMessage() string {
	if len(res.Messages) > 0 {
		return res.Messages[0]
	}
	return ""
}

func NewClientError() GenericResponse {
	return GenericResponse{
		Status:   "Client-Side Error",
		Code:     3,
		Messages: []string{"failed to connect to server"},
		Metadata: map[string]string{},
		Payload:  nil,
	}
}
