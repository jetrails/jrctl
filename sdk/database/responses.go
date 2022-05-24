package database

type ListDatabasesResponse struct {
	Status   string                 `json:"status"`
	Code     int                    `json:"code"`
	Messages []string               `json:"messages"`
	Metadata map[string]interface{} `json:"metadata"`
	Payload  []DatabaseWithUsers    `json:"payload"`
}

type ListUsersResponse struct {
	Status   string                 `json:"status"`
	Code     int                    `json:"code"`
	Messages []string               `json:"messages"`
	Metadata map[string]interface{} `json:"metadata"`
	Payload  []UserWithDatabases    `json:"payload"`
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
	Payload  ConfirmPayload         `json:"payload"`
}

type UserCreateResponse struct {
	Status   string                 `json:"status"`
	Code     int                    `json:"code"`
	Messages []string               `json:"messages"`
	Metadata map[string]interface{} `json:"metadata"`
	Payload  interface{}            `json:"payload"`
}

type UserDeleteResponse struct {
	Status   string                 `json:"status"`
	Code     int                    `json:"code"`
	Messages []string               `json:"messages"`
	Metadata map[string]interface{} `json:"metadata"`
	Payload  ConfirmPayload         `json:"payload"`
}

type UserPasswordResponse struct {
	Status   string                 `json:"status"`
	Code     int                    `json:"code"`
	Messages []string               `json:"messages"`
	Metadata map[string]interface{} `json:"metadata"`
	Payload  ConfirmPayload         `json:"payload"`
}

type LinkResponse struct {
	Status   string                 `json:"status"`
	Code     int                    `json:"code"`
	Messages []string               `json:"messages"`
	Metadata map[string]interface{} `json:"metadata"`
	Payload  interface{}            `json:"payload"`
}

type UnlinkResponse struct {
	Status   string                 `json:"status"`
	Code     int                    `json:"code"`
	Messages []string               `json:"messages"`
	Metadata map[string]interface{} `json:"metadata"`
	Payload  ConfirmPayload         `json:"payload"`
}
