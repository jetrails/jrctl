package database

type ListResponse struct {
	Status   string                 `json:"status"`
	Code     int                    `json:"code"`
	Messages []string               `json:"messages"`
	Metadata map[string]interface{} `json:"metadata"`
	Payload  []Database             `json:"payload"`
}

type CreateResponse struct {
	Status   string                 `json:"status"`
	Code     int                    `json:"code"`
	Messages []string               `json:"messages"`
	Metadata map[string]interface{} `json:"metadata"`
	Payload  interface{}            `json:"payload"`
}

type DeleteResponse struct {
	Status   string                 `json:"status"`
	Code     int                    `json:"code"`
	Messages []string               `json:"messages"`
	Metadata map[string]interface{} `json:"metadata"`
	Payload  string                 `json:"payload"`
}

type UserAddResponse struct {
	Status   string                 `json:"status"`
	Code     int                    `json:"code"`
	Messages []string               `json:"messages"`
	Metadata map[string]interface{} `json:"metadata"`
	Payload  interface{}            `json:"payload"`
}

type UserRemoveResponse struct {
	Status   string                 `json:"status"`
	Code     int                    `json:"code"`
	Messages []string               `json:"messages"`
	Metadata map[string]interface{} `json:"metadata"`
	Payload  string                 `json:"payload"`
}

type UserPasswordResponse struct {
	Status   string                 `json:"status"`
	Code     int                    `json:"code"`
	Messages []string               `json:"messages"`
	Metadata map[string]interface{} `json:"metadata"`
	Payload  string                 `json:"payload"`
}
