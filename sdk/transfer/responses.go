package transfer

type ErrorResponse struct {
	Name string        `json:"name"`
	Message string     `json:"message"`
	Code int           `json:"code"`
	Type string        `json:"type"`
	Data interface {}  `json:"data"`
}

type SendResponse struct {
	Identifier string  `json:"id"`
	Password string    `json:"password"`
	TTL int            `json:"ttl"`
}

type ReceiveResponse struct {
	Bytes [] byte
}
