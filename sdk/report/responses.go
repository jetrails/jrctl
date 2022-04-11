package report

type AuditResponse struct {
	Status   string                 `json:"status"`
	Code     int                    `json:"code"`
	Messages []string               `json:"messages"`
	Metadata map[string]interface{} `json:"metadata"`
	Payload  *AuditData             `json:"payload"`
}
