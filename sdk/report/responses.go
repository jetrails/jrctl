package report

import (
	"github.com/jetrails/jrctl/sdk/api"
)

type AuditResponse struct {
	api.GenericResponse
	Payload *AuditData `json:"payload"`
}
